package handlers

import (
	"encoding/json"
	"net/http"
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/services"
	"sane-discourse-backend/pkg/logger"

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
	log := logger.GetLogger()

	var createPostRequest CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&createPostRequest); err != nil {
		log.WithField("error", err.Error()).Error("CreatePost: Invalid request body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.WithField("input", createPostRequest).Info("CreatePost: Request received")

	post, err := h.postService.CreatePost(createPostRequest.URL)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"input": createPostRequest,
			"error": err.Error(),
		}).Error("CreatePost: Request failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.WithFields(map[string]interface{}{
		"input":  createPostRequest,
		"output": post,
	}).Info("CreatePost: Request successful")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

type AddPostRequest struct {
	// UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Post models.Post `json:"post" bson:"post"`
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	var addPostRequest AddPostRequest
	if err := json.NewDecoder(r.Body).Decode(&addPostRequest); err != nil {
		log.WithField("error", err.Error()).Error("AddPost: Invalid request body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.WithField("input", addPostRequest).Info("AddPost: Request received")

	userIDInterface := r.Context().Value("user_id")
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		http.Error(w, "UserID should be of type primitive.ObjectID", http.StatusInternalServerError)
		return
	}

	post, err := h.postService.AddPost(addPostRequest.Post, userID)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"input": addPostRequest,
			"error": err.Error(),
		}).Error("AddPost: Request failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.WithFields(map[string]interface{}{
		"input":  addPostRequest,
		"output": post,
	}).Info("AddPost: Request successful")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// type GetUserPostsRequest struct {
// 	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
// }

func (h *PostHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

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

	log.WithFields(map[string]interface{}{
		"post_count": len(posts),
	}).Info("GetUserPosts: Request successful")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// type GetUserFeedRequest struct {
// 	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
// }

func (h *PostHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	log.Info("GetFeed: Request received")

	posts, err := h.postService.GetFeed()
	if err != nil {
		log.WithField("error", err.Error()).Error("GetFeed: Request failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.WithField("post_count", len(posts)).Info("GetFeed: Request successful")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
