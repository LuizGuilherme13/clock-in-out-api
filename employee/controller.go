package employee

import (
	"encoding/json"
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/database"
)

func Controller(mux *http.ServeMux) {
	mux.HandleFunc("POST /employee", createEmployee)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	ee := &Model{}

	err := json.NewDecoder(r.Body).Decode(ee)
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

	err = Create(conn, ee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
