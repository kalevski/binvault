package middlewares

import (
	"binvault/pkg/cfg"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lestrrat-go/jwx/jwt"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		auth := cfg.GetAuth()
		if !auth.Enabled {
			next(w, r, ps)
			return
		}

		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if len(token) < 7 || token[:7] != "Bearer " {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token = token[7:]

		log.Println("token %s", token)

		verifiedToken, err := jwt.ParseRequest(r, jwt.WithKeySet(*auth.Jwk))
		if err != nil {
			fmt.Printf("failed to verify token from HTTP request: %s\n", err)
			return
		}
		fmt.Println(verifiedToken)

		next(w, r, ps)
	}
}
