package models

import (
	"time"

	"github.com/LuizGuilherme13/clock-in-api/db"
)

const (
	TimeEntryIn  = "in"
	TimeEntryOut = "out"
)

type TimeEntry struct {
	EmployeeId int64  `json:"employee_id"`
	Type       string `json:"type_entry"` // In or Out
}

func (t *TimeEntry) Punch() error {
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

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return err
	}

	now := time.Now().In(location)

	query := "INSERT INTO time_entries(employee_id, time_entry, type) "
	query += "VALUES($1, $2, $3)"
	_, err = tx.Exec(query, t.EmployeeId, now, t.Type)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (t *TimeEntry) VerifyLastEntryPoint() error {
	conn, err := db.OpenConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := "SELECT t.type FROM time_entries t "
	query += "INNER JOIN employees e"
	query += " ON t.employee_id = e.id "
	query += "WHERE e.id = $1"

	var lastEntryPoint string
	if err := conn.QueryRow(query, t.EmployeeId).Scan(&lastEntryPoint); err != nil {
		return err
	}

	t.Type = TimeEntryIn
	if lastEntryPoint == TimeEntryIn {
		t.Type = TimeEntryOut
	}
	return nil
}
