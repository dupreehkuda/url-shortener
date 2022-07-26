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

	address, isPresent := os.LookupEnv("SERVER_ADDRESS")
	if isPresent != true {
		address = "localhost:8080"
	}

	log.Fatal(server.Run(address))
}
