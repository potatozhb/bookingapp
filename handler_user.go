package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/potatozhb/bookingapp/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parammeters struct {
		Name string `json:"name"`
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

	errdb := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        newid,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if errdb != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user, errget := apiCfg.DB.GetUserByID(r.Context(), newid)
	if errget != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, user)
}
