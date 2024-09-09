package models

import (
	"github.com/jmoiron/sqlx"
)

type Movie struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	ReleaseDate string `json:"release_date" db:"release_date"`
	Duration    int    `json:"duration" db:"duration"`
	TrailerURL  string `json:"trailer_url" db:"trailer_url"`
}

func GetMovies(db *sqlx.DB) ([]Movie, error) {
	var movies []Movie
	err := db.Select(&movies, "SELECT * FROM movie")
	return movies, err
}

func CreateMovie(db *sqlx.DB, m *Movie) error {
	res, err := db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
	if err != nil {
		return err
	}

	m.ID, err = res.LastInsertId()
	return err
}
