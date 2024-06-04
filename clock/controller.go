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

func Controller(mux *http.ServeMux) {
	mux.HandleFunc("POST /clock", recordTime)
	mux.HandleFunc("GET /clock/{id}", getRecords)
}

//	@Summary	Register a clock-in/clock-out time for a employee
//	@ID			register-clock
//	@Tags		clock
//	@Accept		json
//	@Produce	json
//	@Param		employee	body		employee.Model	true	"Payload"
//	@Success	201			{object}	Records
//	@Failure	400			{string}	string	"Bad Request"
//	@Failure	500			{string}	string	"Internal Server Error"
//	@Router		/clock [post]
func recordTime(w http.ResponseWriter, r *http.Request) {
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

	err = RecordTime(conn, foundEE.Id, now.Format(time.DateTime), clockType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clocks, err := GetAllRecords(conn, foundEE.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clocks.Message = fmt.Sprintf("Welcome, %s!", foundEE.UserName)
	if clockType == Out {
		clocks.Message = fmt.Sprintf("See you later, %s!", foundEE.UserName)
	}

	clocks.TotalHoursWorked = calcHoursWorked(clocks)

	err = json.NewEncoder(w).Encode(clocks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getRecords(w http.ResponseWriter, r *http.Request) {

	eeId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
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

	punches, err := GetAllRecords(conn, eeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	punches.TotalHoursWorked = calcHoursWorked(punches)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(punches); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func calcHoursWorked(punches *Records) string {
	var totalWorked time.Duration
	for i := 0; i < len(punches.Clocks); i++ {
		if punches.Clocks[i].Type != Out {
			continue
		}

		in := punches.Clocks[i-1]
		out := punches.Clocks[i]
		totalWorked += out.DateHour.Sub(in.DateHour)
	}

	hours := int(totalWorked.Hours())
	minutes := int(totalWorked.Minutes()) % 60

	return fmt.Sprintf("%d:%02d", hours, minutes)
}
