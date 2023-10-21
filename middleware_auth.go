package main

import (
	"fmt"
	"net/http"

	"github.com/ivanseibel/learning-go-rss-aggregator/internal/auth"
	"github.com/ivanseibel/learning-go-rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Error getting API key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
