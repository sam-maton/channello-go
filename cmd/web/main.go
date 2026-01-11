package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/sam-maton/channello-go/internal/cache"
)

type application struct {
	logger *slog.Logger
	cache  *cache.Cache
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	cacheUrl := flag.String("cache", "./internal/data/channel-cache.json", "Path to channel cache file")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	memCache, err := cache.LoadCache(*cacheUrl)

	if err != nil {
		logger.Error("unable to load cache file", slog.String("error", err.Error()))
		os.Exit(1)
	}

	app := application{
		logger: logger,
		cache:  &memCache,
	}

	logger.Info(fmt.Sprintf("starting server on http://localhost%s/", *addr))

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
