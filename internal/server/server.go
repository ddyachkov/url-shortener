package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	handler "github.com/ddyachkov/url-shortener/internal/handlers"
	"github.com/ddyachkov/url-shortener/internal/storage"
)

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	storage := storage.NewURLStorage()
	storage.LoadData()

	service := app.NewURLShortener(&storage)
	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: handler.NewURLHandler(&service),
	}

	go func() {
		log.Println("Server starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped")
}
