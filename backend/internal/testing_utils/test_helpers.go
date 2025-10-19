package testing

import (
	"context"
	"net/http"
	"sane-discourse-backend/internal/handlers"
	"sane-discourse-backend/internal/middleware"
	"sane-discourse-backend/internal/repositories"
	"sane-discourse-backend/internal/services"
	"sane-discourse-backend/pkg/logger"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupTestServer(t *testing.T) {
	logger.Init()
	log := logger.GetLogger()
	log.Info("Starting sane-discourse-backend server")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")
	}
	log.Info("Successfully connected to MongoDB")

	// Initialize repositories
	userRepo := repositories.NewUserRepository(client)
	postRepo := repositories.NewPostRepository(client)
	reactionRepo := repositories.NewReactionRepository(client)
	userpageRepo := repositories.NewUserpageRepository(client)

	userService := services.NewUserService(userRepo, userpageRepo)
	postService := services.NewPostService(postRepo, userRepo, reactionRepo)
	reactionService := services.NewReactionService(reactionRepo)

	// userHandler := controllers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	reactionHandler := handlers.NewReactionHandler(reactionService)

	authHandler := handlers.NewAuthHandler(userService)

	// Create a new ServeMux
	r := chi.NewRouter()
	// mux := http.NewServeMux()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Register routes
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

	log.Info("Server starting on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.WithError(err).Fatal("Server failed to start")
	}
}
