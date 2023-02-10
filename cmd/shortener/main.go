package main

import (
	"log"
	"net/http"

	"github.com/ddyachkov/url-shortener/internal/app"
	handler "github.com/ddyachkov/url-shortener/internal/handlers"
	"github.com/ddyachkov/url-shortener/internal/storage"
)

func main() {
	storage := storage.NewURLStorage()
	service := app.NewURLShortener(&storage)
	server := http.Server{
		Addr:    ":8080",
		Handler: handler.NewURLHandler(&service),
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
