package main

import (
	"log"
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee", controllers.EmployeeController)

	log.Println("Listening on port :8080...")
	http.ListenAndServe(":8080", mux)
}
