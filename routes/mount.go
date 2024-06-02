package routes

import (
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/employee"
	"github.com/LuizGuilherme13/clock-in-api/toclock"
)

// Mount ...
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/employee", employee.Controller)
	mux.HandleFunc("/punch", toclock.Controller)
}
