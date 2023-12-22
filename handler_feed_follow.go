package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gandresto/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json:", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error following the feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feed_follow))
}
