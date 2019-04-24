package main

import (
	"net/http"
	"os"

	"github.com/swensonhe/fanatick-backend/fanatick/api"
	"github.com/swensonhe/fanatick-backend/fanatick/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		ta := &jwt.TokenAuth{
			Secret:        os.Getenv("CLIENT_SECRET"),
			TokenDuration: 10,
		}

		_, err := ta.Authenticate(token)
		if err != nil {
			err = api.ErrorUnauthorized(`token invalid`)
			NewErrorWriter(w).Write(err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
