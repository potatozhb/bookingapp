package main

import (
	"net/http"

	"github.com/potatozhb/bookingapp/internal/auth"
	"github.com/potatozhb/bookingapp/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		handler(w, r, user)
	}
}
