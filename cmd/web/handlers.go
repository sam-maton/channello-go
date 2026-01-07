package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sam-maton/channello-go/internal/cache"
)

var approveChannels = []string{
	"UC5QwYlOxcT1higtcJVGzCCg",
	"UC6z0E8nSfCvelwA3bon_phg",
}

func getVideoFeed(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request URL:", r.URL.Path)

	cache, err := cache.LoadCache("./internal/data/channel-cache.json")

	_, nonCahced := cache.GetChannelFeeds(approveChannels)

	fmt.Println(nonCahced)

	if err != nil {
		log.Print(err.Error())

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

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

func getVideoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "Video for ID %s", id)
}
