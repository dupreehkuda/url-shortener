package main

import (
	"github.com/dupreehkuda/url-shortener/internal/handlers"
	"github.com/dupreehkuda/url-shortener/internal/server"
	"github.com/dupreehkuda/url-shortener/internal/storage"
	"log"
	"os"
)

func main() {
	storage := storage.New()
	service := handlers.New(storage)

	server := server.New(service).GetRouter()

	address := os.Getenv("SERVER_ADDRESS")
	log.Fatal(server.Run(address))
}
