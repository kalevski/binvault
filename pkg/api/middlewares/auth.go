package middlewares

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/services/auth"
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

		token := helpers.GetRequestToken(r)
		if token == nil {
			helpers.SendError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		_, err := auth.ValidateJWT(keys.PublicKey, *token)
		if err != nil {
			helpers.SendError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		next(w, r, ps)
	}
}
