package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/models"
)

func TimeEntriesController(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, err := employee.Exists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Funcionário não existe ou inválido", http.StatusBadRequest)
		return
	}

	err = punchClock(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func punchClock(employee *models.Employee) error {
	entryController := models.TimeEntry{EmployeeId: employee.Id}

	err := entryController.VerifyLastEntryPoint()
	if err != nil {
		return err
	}

	err = entryController.Punch()
	if err != nil {
		return err
	}

	return nil
}
