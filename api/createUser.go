package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {

	userStruct := User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userStruct); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Problem")
		return
	}

	user, err := cfg.queries.CreateUser(r.Context(), userStruct.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Problem")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
