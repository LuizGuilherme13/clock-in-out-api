package models

import "github.com/LuizGuilherme13/clock-in-api/db"

type Employee struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (e *Employee) Save() error {
	conn, err := db.OpenConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "INSERT INTO employees(username, password) VALUES($1, $2)"
	_, err = tx.Exec(query, e.UserName, e.Password)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
