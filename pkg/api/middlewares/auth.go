package middlewares

import (
	"binvault/pkg/auth"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		keys := auth.GetAuth()
		if !keys.Enabled {
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

		claims, err := auth.ValidateJWT(keys.PublicKey, token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Println("Claims:", claims)
		next(w, r, ps)
	}
}
