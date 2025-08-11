package handlers

import (
	"encoding/json"
	"net/http"
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

type CreatePostsRequest struct {
	URLs []string `json:"urls" bson:"urls"`
}

func dereferencePostSlice(posts []*models.Post) []models.Post {
	result := make([]models.Post, len(posts))

	for i, post := range posts {
		result[i] = *post // Dereference each pointer
	}

	return result
}

func (h *PostHandler) CreatePosts(w http.ResponseWriter, r *http.Request) {
	var createPostsRequest CreatePostsRequest
	if err := json.NewDecoder(r.Body).Decode(&createPostsRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	postReferences, err := h.postService.CreatePosts(createPostsRequest.URLs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	posts := dereferencePostSlice(postReferences)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

type AddPostsRequest struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Posts  []models.Post      `json:"posts" bson:"posts"`
}

func (h *PostHandler) AddPosts(w http.ResponseWriter, r *http.Request) {
	var addPostsRequest AddPostsRequest
	if err := json.NewDecoder(r.Body).Decode(&addPostsRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	posts, err := h.postService.AddPosts(addPostsRequest.Posts, addPostsRequest.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

type GetUserPostsRequest struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
}

func (h *PostHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	var getUserPostsRequest GetUserPostsRequest
	if err := json.NewDecoder(r.Body).Decode(&getUserPostsRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	posts, err := h.postService.GetUserPosts(getUserPostsRequest.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

type GetUserFeedRequest struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
}

func (h *PostHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	/*var getUserFeedRequest GetUserFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&getUserFeedRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}*/
	posts, err := h.postService.GetFeed()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
