package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/{$}", app.getVideoFeed)
	mux.HandleFunc("/feed", app.getVideoFeed)
	mux.HandleFunc("GET /videos/{id}", app.getVideoById)

	return mux
}
