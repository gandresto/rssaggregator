package main

import (
	"fmt"
	"net/http"

	"github.com/gandresto/rssaggregator/internal/auth"
	"github.com/gandresto/rssaggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprint("Error parsing json:", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprint("Could not get the user:", err))
			return
		}

		handler(w, r, user)
	}
}
