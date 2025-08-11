package main

import (
	"context"
	"net/http"

	"sane-discourse-backend/internal/handlers"
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
	postService := services.NewPostService(postRepo)
	reactionService := services.NewReactionService(reactionRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	reactionHandler := handlers.NewReactionHandler(reactionService)

	http.HandleFunc("/user/login", userHandler.LoginUser)

	_ = postHandler
	_ = reactionHandler

	http.ListenAndServe(":3000", nil)
}
