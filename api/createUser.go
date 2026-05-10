package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {

	user := User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Problem")
	}

}
