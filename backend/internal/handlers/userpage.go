package handlers

import (
	"encoding/json"
	"net/http"
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserpageHandler struct {
	userpageService *services.UserpageService
}

func NewUserpageHandler(s *services.UserpageService) *UserpageHandler {
	return &UserpageHandler{
		userpageService: s,
	}
}

type AddComponentRequest struct {
	Index     int              `json:"index"`
	Component models.Component `json:"component"`
}

func (h *UserpageHandler) AddComponent(w http.ResponseWriter, r *http.Request) {
	var addComponentRequest AddComponentRequest
	json.NewDecoder(r.Body).Decode(&addComponentRequest)

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	userpage, err := h.userpageService.AddComponent(
		userID,
		addComponentRequest.Index,
		&addComponentRequest.Component,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(*userpage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

type MoveComponentRequest struct {
	PrevIndex int `json:"prev_index" bson:"prev_index"`
	NewIndex  int `json:"new_index" bson:"new_index"`
}

func (h *UserpageHandler) MoveComponent(w http.ResponseWriter, r *http.Request) {
	var moveComponentRequest MoveComponentRequest
	json.NewDecoder(r.Body).Decode(&moveComponentRequest)

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	userpage, err := h.userpageService.MoveComponent(
		userID,
		moveComponentRequest.PrevIndex,
		moveComponentRequest.NewIndex,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(*userpage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
