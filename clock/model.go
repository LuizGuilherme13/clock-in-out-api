package clock

import "time"

type Clock int

const (
	In Clock = iota + 1
	Out
)

type Model struct {
	EmployeeId int64  `json:"employee_id"`
	Type       string `json:"type_entry"` // In or Out
}

type ClockPunch struct {
	EmployeeId   int64     `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	TimeEntry    time.Time `json:"date_hour"`
	Type         Clock     `json:"type"` // In or Out
}

type Punches struct {
	Clocks           []ClockPunch
	TotalHoursWorked string
}
