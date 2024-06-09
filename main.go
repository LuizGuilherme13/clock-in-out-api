package main

import (
	"log"
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/database"
	"github.com/LuizGuilherme13/clock-in-api/routes"
)

func init() {
	database.InitTables()
}

// @title			Clock-In/Out API
// @version		1.0
// @description	API to record employees' clock-in and clock-out times.
func main() {
	mux := http.NewServeMux()
	routes.Mount(mux)

	log.Println("Listening on port :8080...")
	http.ListenAndServe(":8080", mux)
}
