package clock

import (
	"database/sql"
)

func RecordTime(conn *sql.DB, employeeId int64, dateHour string, clockType TypeRecord) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "INSERT INTO clock_records(employee_id, date_hour, type)"
	query += "VALUES($1, $2, $3)"

	_, err = tx.Exec(query, employeeId, dateHour, clockType)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ArrivingOrLeaving(conn *sql.DB, employeeID int64, date string) (TypeRecord, error) {
	query := "SELECT type FROM clock_records "
	query += "WHERE"
	query += "  employee_id = $1"
	query += "  AND date_hour::date = $2 "
	query += "ORDER BY id DESC "
	query += "LIMIT 1"

	var lastClock int
	err := conn.QueryRow(query, employeeID, date).Scan(&lastClock)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if lastClock == 0 || TypeRecord(lastClock) == Out {
		return In, nil
	}

	return Out, nil
}

func GetAllRecords(conn *sql.DB, employeeId int64) (*Records, error) {
	query := "SELECT ee.id, ee.username, cr.date_hour, cr.type "
	query += "FROM clock_records cr "
	query += "INNER JOIN employees ee"
	query += "  ON ee.id = cr.employee_id "
	query += "WHERE ee.id = $1"

	rows, err := conn.Query(query, employeeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p := Records{}
	cp := Record{}
	t := int(0)

	for rows.Next() {
		err = rows.Scan(&cp.EmployeeId, &cp.EmployeeName, &cp.DateHour, &t)
		if err != nil {
			return nil, err
		}
		cp.Type = TypeRecord(t)

		p.Clocks = append(p.Clocks, cp)
	}

	return &p, nil
}
