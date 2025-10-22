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

func (h *UserpageHandler) GetUserpage(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	userpage, err := h.userpageService.GetUserpage(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(*userpage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type UpdateComponentRequest struct {
	Index     int              `json:"index"`
	Component models.Component `json:"component"`
}

func (h *UserpageHandler) UpdateComponent(w http.ResponseWriter, r *http.Request) {
	var updateComponentRequest UpdateComponentRequest
	json.NewDecoder(r.Body).Decode(&updateComponentRequest)

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	userpage, err := h.userpageService.UpdateComponent(
		userID,
		updateComponentRequest.Index,
		&updateComponentRequest.Component,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(*userpage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type DeleteComponentRequest struct {
	Index int `json:"index"`
}

func (h *UserpageHandler) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	var deleteComponentRequest DeleteComponentRequest
	json.NewDecoder(r.Body).Decode(&deleteComponentRequest)

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	userpage, err := h.userpageService.DeleteComponent(
		userID,
		deleteComponentRequest.Index,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(*userpage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(*userpage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
