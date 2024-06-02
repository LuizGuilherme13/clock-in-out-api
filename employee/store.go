package employee

import (
	"database/sql"
)

func Create(conn *sql.DB, employee *Model) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "INSERT INTO employees(username, password) "
	query += "VALUES($1, $2) "
	query += "RETURNING id "

	var id int64
	err = tx.QueryRow(query, employee.UserName, employee.Password).Scan(&id)
	if err != nil {
		return err
	}

	employee.Id = id

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func FindByName(conn *sql.DB, name string) (*Model, error) {
	query := "SELECT ee.id, ee.username, ee.password FROM employees ee "
	query += "WHERE ee.username = $1"

	ee := Model{}
	err := conn.QueryRow(query, name).Scan(&ee.Id, &ee.UserName, &ee.Password)
	if err != nil {
		return nil, err
	}

	return &ee, nil
}
