package main

import (
	"github.com/dupreehkuda/url-shortener/internal/handlers"
	"github.com/dupreehkuda/url-shortener/internal/server"
	"github.com/dupreehkuda/url-shortener/internal/storage"
)

func main() {
	storage := storage.New()
	service := handlers.New(storage)

	server := server.New(service)
	server.Launch()
}
