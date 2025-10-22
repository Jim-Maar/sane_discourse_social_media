package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"sane-discourse-backend/internal/auth"
	"sane-discourse-backend/internal/handlers"
	"sane-discourse-backend/internal/middleware"
	"sane-discourse-backend/internal/repositories"
	"sane-discourse-backend/internal/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	gob.Register(primitive.ObjectID{})
	auth.NewAuth()

	mongoURI := "mongodb://admin:dev_admin_password@localhost:27017"
	mongoUsername := "admin"
	mongoPassword := "dev_admin_password"
	port := "3000"

	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetAuth(options.Credential{
			Username: mongoUsername,
			Password: mongoPassword,
		})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	userRepo := repositories.NewUserRepository(client)
	postRepo := repositories.NewPostRepository(client)
	reactionRepo := repositories.NewReactionRepository(client)
	userpageRepo := repositories.NewUserpageRepository(client)

	userService := services.NewUserService(userRepo, userpageRepo)
	postService := services.NewPostService(postRepo, userRepo, reactionRepo)
	reactionService := services.NewReactionService(reactionRepo)
	userpageService := services.NewUserpageService(userpageRepo)

	// userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	reactionHandler := handlers.NewReactionHandler(reactionService)
	authHandler := handlers.NewAuthHandler(userService)
	userpageHandler := handlers.NewUserpageHandler(userpageService)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// r.Post("/user/login", userHandler.LoginUser)
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleWare)
		r.Put("/user/posts/create", postHandler.CreatePost)
		r.Put("/user/posts/add", postHandler.AddPost)
		r.Get("/user/posts", postHandler.GetUserPosts)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleWare)
		r.Put("/auth/me", authHandler.GetCurrentUser)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleWare)
		r.Get("/userpage", userpageHandler.GetUserpage)
		r.Put("/userpage/component/add", userpageHandler.AddComponent)
		r.Put("/userpage/component/update", userpageHandler.UpdateComponent)
		r.Delete("/userpage/component/delete", userpageHandler.DeleteComponent)
		r.Put("/userpage/component/move", userpageHandler.MoveComponent)
	})
	r.Get("/home", postHandler.GetFeed)

	r.Get("/auth/{provider}", authHandler.BeginAuthProviderCallback)
	// r.Get("/logout/{provider}", authHandler.GetLogoutFunction)
	r.Get("/auth/{provider}/callback", authHandler.GetAuthCallbackFunction)

	_ = reactionHandler

	serverAddr := fmt.Sprintf(":%s", port)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
