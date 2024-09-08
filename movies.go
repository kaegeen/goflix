package handlers

import (
	"encoding/json"
	"net/http"

	"goflix/internal/models"

	"github.com/jmoiron/sqlx"
)

type MovieHandler struct {
	db        *sqlx.DB
	jwtSecret string
}

func NewMovieHandler(db *sqlx.DB, jwtSecret string) *MovieHandler {
	return &MovieHandler{db: db, jwtSecret: jwtSecret}
}

func (h *MovieHandler) HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := models.GetMovies(h.db)
	if err != nil {
		http.Error(w, "Could not fetch movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) HandleCreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := models.CreateMovie(h.db, &movie); err != nil {
		http.Error(w, "Could not create movie", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Handle JWT login
}
