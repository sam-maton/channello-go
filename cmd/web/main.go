package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", getVideoFeed)
	mux.HandleFunc("/feed", getVideoFeed)
	mux.HandleFunc("GET /videos/{id}", getVideoById)

	logger.Info(fmt.Sprintf("starting server on http://localhost%s/", *addr))

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
