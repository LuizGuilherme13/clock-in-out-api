package main

import (
	"log"
	"net/http"
)

//	@title			Clock-In/Out API
//	@version		1.0
//	@description	API to record employees' clock-in and clock-out times.
func main() {
	mux := http.NewServeMux()
	MountRoutes(mux)

	log.Println("Listening on port :8080...")
	http.ListenAndServe(":8080", mux)
}
