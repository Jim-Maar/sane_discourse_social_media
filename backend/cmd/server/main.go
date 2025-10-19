package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"sane-discourse-backend/internal/auth"
	"sane-discourse-backend/internal/handlers"
	"sane-discourse-backend/internal/middleware"
	"sane-discourse-backend/internal/repositories"
	"sane-discourse-backend/internal/services"
	"sane-discourse-backend/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger.Init()
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatal("Error loading .env file")
	}

	auth.NewAuth()

	mongoURI := os.Getenv("MONGODB_URI")
	mongoUsername := os.Getenv("MONGODB_USERNAME")
	mongoPassword := os.Getenv("MONGODB_PASSWORD")
	port := os.Getenv("PORT")

	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}
	if mongoUsername == "" {
		log.Fatal("MONGODB_USERNAME environment variable is not set")
	}
	if mongoPassword == "" {
		log.Fatal("MONGODB_PASSWORD environment variable is not set")
	}
	if port == "" {
		port = "3000" // Default port if not specified
		log.Warn("PORT environment variable not set, using default: 3000")
	}

	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetAuth(options.Credential{
			Username: mongoUsername,
			Password: mongoPassword,
		})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.WithError(err).Fatal("Failed to ping MongoDB")
	}

	userRepo := repositories.NewUserRepository(client)
	postRepo := repositories.NewPostRepository(client)
	reactionRepo := repositories.NewReactionRepository(client)
	userpageRepo := repositories.NewUserpageRepository(client)

	userService := services.NewUserService(userRepo, userpageRepo)
	postService := services.NewPostService(postRepo, userRepo, reactionRepo)
	reactionService := services.NewReactionService(reactionRepo)

	// userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	reactionHandler := handlers.NewReactionHandler(reactionService)

	authHandler := handlers.NewAuthHandler(userService)

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
	r.Get("/home", postHandler.GetFeed)

	r.Get("/auth/{provider}", authHandler.BeginAuthProviderCallback)
	// r.Get("/logout/{provider}", authHandler.GetLogoutFunction)
	r.Get("/auth/{provider}/callback", authHandler.GetAuthCallbackFunction)

	_ = reactionHandler

	serverAddr := fmt.Sprintf(":%s", port)
	log.Infof("Server starting on port %s", port)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.WithError(err).Fatal("Server failed to start")
	}
}
