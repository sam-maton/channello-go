package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := application{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", app.getVideoFeed)
	mux.HandleFunc("/feed", app.getVideoFeed)
	mux.HandleFunc("GET /videos/{id}", app.getVideoById)

	logger.Info(fmt.Sprintf("starting server on http://localhost%s/", *addr))

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
