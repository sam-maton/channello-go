package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var approveChannels = []string{
	"UC5QwYlOxcT1higtcJVGzCCg",
	"UC6z0E8nSfCvelwA3bon_phg",
}

func (app *application) getVideoFeed(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/video-feed.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (app *application) getVideoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "Video for ID %s", id)
}
