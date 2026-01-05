package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Channello!"))
}

func getVideoFeed(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video feed returned"))
}

func getVideoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "Video for ID %s", id)
}
