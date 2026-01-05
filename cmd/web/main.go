package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getVideoFeed)
	mux.HandleFunc("/feed", getVideoFeed)
	mux.HandleFunc("GET /videos/{id}", getVideoById)

	log.Print(("starting server on http://localhost:4000/"))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
