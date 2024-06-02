package main

import (
	"log"
	"net/http"

	"github.com/LuizGuilherme13/clock-in-api/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Mount(mux)

	log.Println("Listening on port :8080...")
	http.ListenAndServe(":8080", mux)
}
