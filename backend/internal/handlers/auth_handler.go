package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sane-discourse-backend/internal/services"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) GetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Printf("Auth error: %v\n", err)
		http.Error(w, fmt.Sprintf("Authentication failed: %v", err), http.StatusInternalServerError)
		return
	}

	dbUser, err := h.userService.LoginUser(user.Name, user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := gothic.Store.Get(r, "auth-session")
	session.Values["user_id"] = dbUser.ID
	session.Save(r, w)

	http.Redirect(w, r, "http://localhost:5173/home", http.StatusFound)
}

func (h *AuthHandler) BeginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.BeginAuthHandler(w, r)
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}
	user, err := h.userService.GetCurrentUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*user)
}
