package main

import (
	"log"
	"net/http"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	handler "github.com/ddyachkov/url-shortener/internal/handlers"
	"github.com/ddyachkov/url-shortener/internal/storage"
)

func main() {
	config := config.GetConfig()
	storage := storage.NewURLStorage()
	service := app.NewURLShortener(&storage)
	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: handler.NewURLHandler(&service, &config),
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
