package middleware

import (
	"context"
	"net/http"

	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := gothic.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userID, ok := session.Values["user_id"]
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		userID = userID.(primitive.ObjectID)
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
