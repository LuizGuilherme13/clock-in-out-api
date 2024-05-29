package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	dataSourceName := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
