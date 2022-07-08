package main

import (
	"github.com/dupreehkuda/url-shortener/cmd/handlers"
	"github.com/dupreehkuda/url-shortener/cmd/storage"
	"log"
	"net/http"
)

func main() {
	storage := storage.New()

	http.HandleFunc("/", handlers.ApiResponse(storage.Shortened))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
