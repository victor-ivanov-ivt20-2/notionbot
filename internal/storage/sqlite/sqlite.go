package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	stmt, err := db.Prepare(
		`
		CREATE TABLE IF NOT EXISTS user(
			id INTEGER PRIMARY KEY
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return &Storage{db: db}, nil
}
