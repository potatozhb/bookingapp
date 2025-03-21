package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/potatozhb/bookingapp/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parammeters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parammeters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	fmt.Println(params.Name)

	newid := uuid.New().String()

	errdb := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        newid,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if errdb != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	feed, errget := apiCfg.DB.GetFeedByID(r.Context(), newid)
	if errget != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, dbFeedsToFeeds(feeds))
}
