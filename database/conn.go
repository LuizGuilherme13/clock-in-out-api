package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	dataSourceName := `
		host=clock-in-postgres-1 
		port=5432 
		user=postgres 
		password=postgres 
		dbname=postgres 
		sslmode=disable
	`

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InitTables() error {
	conn, err := OpenConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS employees(
			id INT GENERATED ALWAYS AS IDENTITY,
			username VARCHAR(50) NOT NULL,
			password VARCHAR(50) NOT NULL,
			PRIMARY KEY(id)
		);`,

		`CREATE TABLE IF NOT EXISTS clock_records(
			id INT GENERATED ALWAYS AS IDENTITY,
			employee_id INT NOT NULL,
			date_hour TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			type SMALLINT NOT NULL,
			PRIMARY KEY(id),
			CONSTRAINT fk_employee
      			FOREIGN KEY(employee_id) 
        			REFERENCES employees(id)
		);`,
	}

	for _, query := range queries {
		_, err := conn.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
