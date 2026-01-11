package main

import (
	"fmt"
	"html/template"
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
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

}

func (app *application) getVideoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "Video for ID %s", id)
}
