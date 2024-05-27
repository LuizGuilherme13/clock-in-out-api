package controllers

import "net/http"

func EmployeeController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		create(w)
	case http.MethodPut:
		update()
	case http.MethodDelete:
		delete()
	default:
		http.Error(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func create(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}
func update() {}
func delete() {}
