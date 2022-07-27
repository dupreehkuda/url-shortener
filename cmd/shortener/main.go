package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/dupreehkuda/url-shortener/internal/handlers"
	"github.com/dupreehkuda/url-shortener/internal/server"
	"github.com/dupreehkuda/url-shortener/internal/storage/filestore"
	"github.com/dupreehkuda/url-shortener/internal/storage/memstore"
	"log"
)

var cfg struct {
	Addr            string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

func main() {

	var err = env.Parse(&cfg)
	if err != nil {
		log.Println(err)
	}

	var storage handlers.Storer

	if cfg.FileStoragePath != "" {
		storage = filestore.New(cfg.FileStoragePath)
		log.Printf("Launched with file")
	} else {
		storage = memstore.New()
		log.Printf("Launched with memory")
	}

	service := handlers.New(storage)
	server := server.New(service)

	server.Launch()
}
