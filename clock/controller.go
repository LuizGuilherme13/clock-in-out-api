package clock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/LuizGuilherme13/clock-in-api/database"
	"github.com/LuizGuilherme13/clock-in-api/employee"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		punch(w, r)
	case http.MethodGet:
		getClocks(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func punch(w http.ResponseWriter, r *http.Request) {
	ee := &employee.Model{}

	err := json.NewDecoder(r.Body).Decode(ee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := database.OpenConn()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer conn.Close()

	foundEE, err := employee.FindByName(conn, ee.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now()

	clockType, err := ArrivingOrLeaving(conn, foundEE.Id, now.Format(time.DateOnly))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Punch(conn, foundEE.Id, now.Format(time.DateTime), clockType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getClocks(w http.ResponseWriter, r *http.Request) {

	eeId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := database.OpenConn()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	punches, err := GetAllPunchClocks(conn, eeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// var intervals  []time.Duration
	var totalWorked time.Duration
	for i := 0; i < len(punches.Clocks); i++ {
		if punches.Clocks[i].Type != Out {
			continue
		}

		in := punches.Clocks[i-1]
		out := punches.Clocks[i]
		totalWorked += out.TimeEntry.Sub(in.TimeEntry)
	}

	hours := int(totalWorked.Hours())
	minutes := int(totalWorked.Minutes()) % 60
	punches.TotalHoursWorked = fmt.Sprintf("%d:%02d", hours, minutes)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(punches); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
