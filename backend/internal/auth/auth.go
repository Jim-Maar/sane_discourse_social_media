package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	key := []byte(jwtSecret)
	if len(key) < 32 {
		log.Fatal("JWT_SECRET must be at least 32 bytes long for AES-256 encryption")
	}
	// if len(key) > 32 {
	// 	key = key[:32]
	// }

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if googleClientId == "" {
		log.Fatal("GOOGLE_CLIENT_ID environment variable is not set")
	}
	if googleClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_SECRET environment variable is not set")
	}

	store := sessions.NewCookieStore(key)
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd
	store.Options.SameSite = http.SameSiteDefaultMode

	gothic.Store = store
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, "http://localhost:3000/auth/google/callback", "email", "profile"),
	)
}
