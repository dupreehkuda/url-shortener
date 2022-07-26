package server

import (
	"github.com/gin-gonic/gin"
	"log"
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

func (s *server) GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/:id", s.handlers.GetShortened())
	r.POST("/", s.handlers.PostShorten())
	r.POST("/api/shorten", s.handlers.ShortenJSON())

	log.Printf("Server created")
	return r
}
