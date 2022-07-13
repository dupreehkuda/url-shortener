package main

import (
	"github.com/dupreehkuda/url-shortener/cmd/handlers"
	"github.com/dupreehkuda/url-shortener/cmd/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	storage := storage.New()

	r := gin.Default()

	r.GET("/:id", handlers.GetShortened(storage.Shortened))
	r.POST("/", handlers.PostShorten(storage.Shortened))

	r.Run()
}
