package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Handlers interface {
	GetShortened() gin.HandlerFunc
	PostShorten() gin.HandlerFunc
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

	log.Printf("Server created")
	return r
}
