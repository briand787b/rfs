package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/render"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

const jwtSecretEnvVar = "JWT_SECRET"

// extract this as well as the init
var tokenAuth *jwtauth.JWTAuth

func setSecret() {
	secret := os.Getenv(jwtSecretEnvVar)
	if secret == "" {
		log.Printf("WARNING: env var %s is empty string\n", jwtSecretEnvVar)
	}

	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func setToken(w http.ResponseWriter) error {
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{
		"user_id": 123,
		"exp":     time.Now().Add(60 * time.Minute).Unix(),
	})

	if err != nil {
		return errors.Wrap(err, "could not encode jwt")
	}

	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	w.Header().Set("Authorization", "BEARER: "+tokenString)
	return nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if err := setToken(w); err != nil {
		// obviously this is not right and should be changed in the future
		render.Render(w, r, ErrNotFound)
	}
}
