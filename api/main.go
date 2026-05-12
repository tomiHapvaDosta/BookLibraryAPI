package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tomiHapvaDosta/BookLibraryAPI/internal/database"
)

type apiConfig struct {
	queries *database.Queries
}

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Print("Problem")
		return
	}

	dbQueries := database.New(db)

	const port = "8080"

	cfg := apiConfig{
		queries: dbQueries,
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.HandleFunc("POST /api/users", cfg.createUser)

	log.Fatal(server.ListenAndServe())
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
