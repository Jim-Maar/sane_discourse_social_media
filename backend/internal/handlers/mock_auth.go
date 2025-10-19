package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sane-discourse-backend/internal/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockAuthHandler struct {
	userService *services.UserService
	userStore   map[string]string
}

func NewMockAuthHandler(userService *services.UserService) *MockAuthHandler {
	return &MockAuthHandler{
		userService: userService,
		userStore:   make(map[string]string),
	}
}

type LoginUserRequest struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

func (h *MockAuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUserRequest LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&loginUserRequest); err != nil {
		log.Printf("LoginUser: Invalid request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user, err := h.userService.LoginUser(loginUserRequest.Name, loginUserRequest.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.userStore["user_id"] = user.ID.Hex()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*user)
}

func (h *MockAuthHandler) MockAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := primitive.ObjectIDFromHex(h.userStore["user_id"])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
