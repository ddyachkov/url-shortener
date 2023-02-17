package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/handler"
	"github.com/ddyachkov/url-shortener/internal/storage"
)

func main() {
	flag.Parse()
	cfg := config.NewServerConfig()
	log.Printf("config: %+v\n", cfg)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	storage := storage.NewURLFileStorage(&cfg)
	storage.LoadData()

	service := app.NewURLShortener(&storage)
	server := http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler.NewURLHandler(&service, &cfg),
	}

	go func() {
		log.Println("server starting...")
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
	log.Println("server stopped")
}
