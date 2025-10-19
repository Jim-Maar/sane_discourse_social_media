package handlers

import (
	"encoding/json"
	"log"
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

type CreatePostRequest struct {
	URL string `json:"url" bson:"url"`
}

func dereferencePostSlice(posts []*models.Post) []models.Post {
	result := make([]models.Post, len(posts))

	for i, post := range posts {
		result[i] = *post // Dereference each pointer
	}

	return result
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var createPostRequest CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&createPostRequest); err != nil {
		log.Printf("CreatePost: Invalid request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	post, err := h.postService.CreatePost(createPostRequest.URL)
	if err != nil {
		log.Printf("CreatePost: Request failed for input %+v: %v", createPostRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

type AddPostRequest struct {
	// UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Post models.Post `json:"post" bson:"post"`
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	var addPostRequest AddPostRequest
	if err := json.NewDecoder(r.Body).Decode(&addPostRequest); err != nil {
		log.Printf("AddPost: Invalid request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	post, err := h.postService.AddPost(addPostRequest.Post, userID)
	if err != nil {
		log.Printf("AddPost: Request failed for input %+v: %v", addPostRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// type GetUserPostsRequest struct {
// 	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
// }

func (h *PostHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// var getUserPostsRequest GetUserPostsRequest
	// if err := json.NewDecoder(r.Body).Decode(&getUserPostsRequest); err != nil {
	// 	log.WithField("error", err.Error()).Error("GetUserPosts: Invalid request body")
	// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 	return
	// }

	// log.WithField("input", getUserPostsRequest).Info("GetUserPosts: Request received")

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	posts, err := h.postService.GetUserPosts(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// type GetUserFeedRequest struct {
// 	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
// }

func (h *PostHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetFeed()
	if err != nil {
		log.Printf("GetFeed: Request failed: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
