package clock

import "time"

type TypeRecord int

const (
	In TypeRecord = iota + 1
	Out
)

type Model struct {
	EmployeeId int64  `json:"employee_id"`
	Type       string `json:"type_entry"` // In or Out
}

type Record struct {
	EmployeeId   int64      `json:"employee_id"`
	EmployeeName string     `json:"employee_name"`
	DateHour     time.Time  `json:"date_hour"`
	Type         TypeRecord `json:"type"` // In or Out
}

type Records struct {
	Message          string   `json:"message,omitempty"`
	Clocks           []Record `json:"records"`
	TotalHoursWorked string   `json:"hours_worked"`
}
