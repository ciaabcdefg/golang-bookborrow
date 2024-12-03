package db

import (
	"borrow/internal/env"
	"database/sql"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	env.Init()

	db, err := sql.Open("postgres", env.DBConnection)
	if err != nil {
		return nil, err
	}
	db.Ping()

	return db, nil
}
