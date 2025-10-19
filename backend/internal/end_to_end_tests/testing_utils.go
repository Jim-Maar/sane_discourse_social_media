package end_to_end_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"sane-discourse-backend/internal/handlers"
	"sane-discourse-backend/internal/repositories"
	"sane-discourse-backend/internal/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupTestServer() *chi.Mux {
	mongoURI := "mongodb://admin:dev_admin_password@localhost:27017"
	mongoUsername := "admin"
	mongoPassword := "dev_admin_password"

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
	mockAuthHander := handlers.NewMockAuthHandler(userService)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Put("/auth/login", mockAuthHander.LoginUser)
	r.Group(func(r chi.Router) {
		r.Use(mockAuthHander.MockAuthMiddleWare)
		r.Put("/user/posts/create", postHandler.CreatePost)
		r.Put("/user/posts/add", postHandler.AddPost)
		r.Get("/user/posts", postHandler.GetUserPosts)
	})
	r.Group(func(r chi.Router) {
		r.Use(mockAuthHander.MockAuthMiddleWare)
		r.Put("/auth/me", authHandler.GetCurrentUser)
	})
	r.Group(func(r chi.Router) {
		r.Use(mockAuthHander.MockAuthMiddleWare)
		r.Put("/userpage/component/add", userpageHandler.AddComponent)
		r.Put("/userpage/component/move", userpageHandler.MoveComponent)
	})
	r.Get("/home", postHandler.GetFeed)

	_ = reactionHandler

	return r
}

func PerformRequest(r http.Handler, method, path string, body any) *httptest.ResponseRecorder {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	return w
}
