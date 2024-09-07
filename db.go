package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(dbPath string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS movie (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		release_date TEXT,
		duration INTEGER,
		trailer_url TEXT
	);

	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);
	`
	db.MustExec(schema)
	return db, nil
}
