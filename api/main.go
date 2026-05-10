package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/tomiHapvaDosta/BookLibraryAPI/internal/database"
)

type apiConfig struct {
	queries *database.Queries
}

func main() {

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Print("Problem")
		return
	}

	dbQueries := database.New(db)

	cfg := apiConfig{
		queries: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users", cfg.createUser)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	resp := struct {
		Error string `json:"error"`
	}{Error: msg}
	respondWithJSON(w, code, resp)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
