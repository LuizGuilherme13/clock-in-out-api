package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/models"
)

func EmployeeController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		create(w, r)
	case http.MethodPut:
		update()
	case http.MethodDelete:
		delete()
	default:
		http.Error(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	newEmployee := models.Employee{}
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := newEmployee.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func update() {}
func delete() {}
