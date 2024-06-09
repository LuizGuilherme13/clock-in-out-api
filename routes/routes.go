package routes

import (
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/clock"
	_ "github.com/LuizGuilherme13/clock-in-api/docs"
	"github.com/LuizGuilherme13/clock-in-api/employee"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/swagger/*", httpSwagger.WrapHandler)

	employee.Controller(mux)
	clock.Controller(mux)
}
