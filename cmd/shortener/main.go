package main

import (
	"github.com/dupreehkuda/url-shortener/internal/handlers"
	"github.com/dupreehkuda/url-shortener/internal/server"
	"github.com/dupreehkuda/url-shortener/internal/storage"
)

func main() {
	storage := storage.New()

	r := server.GetRouter()

	r.GET("/:id", handlers.GetShortened(storage.Shortened))
	r.POST("/", handlers.PostShorten(storage.Shortened))

	r.Run()
}
