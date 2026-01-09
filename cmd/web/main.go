package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", getVideoFeed)
	mux.HandleFunc("/feed", getVideoFeed)
	mux.HandleFunc("GET /videos/{id}", getVideoById)

	log.Printf("starting server on http://localhost%s/", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
