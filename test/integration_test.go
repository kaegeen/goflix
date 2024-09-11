package main

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

type dbStore struct {
	db *sqlx.DB
}

func setupTestDB() (*sqlx.DB, error) {

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE movie (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        release_date TEXT,
        duration INTEGER,
        trailer_url TEXT
    );`

	db.MustExec(schema)

	log.Println("Test database setup complete.")
	return db, nil
}

func (store *dbStore) CreateMovie(m *Movie) error {
	result, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
	if err != nil {
		return err
	}
	m.ID, err = result.LastInsertId()
	return err
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	return movies, err
}

func (store *dbStore) GetMovieById(id int64) (*Movie, error) {
	var movie Movie
	err := store.db.Get(&movie, "SELECT * FROM movie WHERE id=?", id)
	return &movie, err
}

func TestMovieCreate(t *testing.T) {

	db, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	store := &dbStore{db: db}

	movie := &Movie{
		Title:       "Inception",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerURL:  "http://url",
	}

	err = store.CreateMovie(movie)
	assert.NoError(t, err)

	assert.NotZero(t, movie.ID)
}

func TestGetMovies(t *testing.T) {

	db, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	store := &dbStore{db: db}

	movie := &Movie{
		Title:       "Inception",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerURL:  "http://url",
	}
	err = store.CreateMovie(movie)
	assert.NoError(t, err)

	movies, err := store.GetMovies()
	assert.NoError(t, err)
	assert.NotEmpty(t, movies)

	assert.Equal(t, "Inception", movies[0].Title)
}

func TestGetMovieById(t *testing.T) {

	db, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	store := &dbStore{db: db}

	movie := &Movie{
		Title:       "Inception",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerURL:  "http://url",
	}
	err = store.CreateMovie(movie)
	assert.NoError(t, err)

	fetchedMovie, err := store.GetMovieById(movie.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedMovie)

	assert.Equal(t, movie.Title, fetchedMovie.Title)
}

func main() {

	log.Println("Running TestMovieCreate...")
	TestMovieCreate(&testing.T{})

	log.Println("Running TestGetMovies...")
	TestGetMovies(&testing.T{})

	log.Println("Running TestGetMovieById...")
	TestGetMovieById(&testing.T{})
}
