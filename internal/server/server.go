package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Handlers interface {
	GetShortened() gin.HandlerFunc
	PostShorten() gin.HandlerFunc
	ShortenJSON() gin.HandlerFunc
}

type server struct {
	handlers Handlers
}

func New(handlers Handlers) *server {
	return &server{handlers: handlers}
}

func (s *server) Launch() {
	r := gin.Default()

	address, isPresent := os.LookupEnv("SERVER_ADDRESS")
	if !isPresent {
		address = "localhost:8080"
	}

	r.GET("/:id", s.handlers.GetShortened())
	r.POST("/", s.handlers.PostShorten())
	r.POST("/api/shorten", s.handlers.ShortenJSON())

	log.Printf("Server created")
	log.Fatal(r.Run(address))
}
