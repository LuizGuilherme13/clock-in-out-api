package main

import (
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/clock"
	"github.com/LuizGuilherme13/clock-in-api/employee"
)

func MountRoutes(mux *http.ServeMux) {
	employee.Controller(mux)
	clock.Controller(mux)
}
