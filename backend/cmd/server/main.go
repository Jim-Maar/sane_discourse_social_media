package main

import (
	"context"
	"net/http"

	"sane-discourse-backend/internal/handlers"
	"sane-discourse-backend/internal/middleware"
	"sane-discourse-backend/internal/repositories"
	"sane-discourse-backend/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(client)
	postRepo := repositories.NewPostRepository(client)
	reactionRepo := repositories.NewReactionRepository(client)

	// Initialize services
	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo, userRepo, reactionRepo)
	reactionService := services.NewReactionService(reactionRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	reactionHandler := handlers.NewReactionHandler(reactionService)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/user/login", userHandler.LoginUser)
	mux.HandleFunc("/user/posts/create", postHandler.CreatePosts)
	mux.HandleFunc("/user/posts/add", postHandler.AddPost)
	mux.HandleFunc("/user/posts", postHandler.GetUserPosts)
	mux.HandleFunc("/home", postHandler.GetFeed)

	_ = reactionHandler

	// Wrap with CORS middleware
	handler := middleware.CORSMiddleware(mux)

	http.ListenAndServe(":3000", handler)
}
