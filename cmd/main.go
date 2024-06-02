package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	MountRoutes(mux)

	log.Println("Listening on port :8080...")
	http.ListenAndServe(":8080", mux)
}
