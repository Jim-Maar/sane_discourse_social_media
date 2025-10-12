package middleware

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestingAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := primitive.NewObjectID()
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
