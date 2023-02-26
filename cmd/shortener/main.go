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
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	flag.Parse()
	cfg := config.NewServerConfig()
	log.Printf("config: %+v\n", cfg)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var service app.URLShortener
	var dbpool *pgxpool.Pool
	switch {
	case cfg.DatabaseDsn != "":
		var err error
		dbpool, err = pgxpool.New(ctx, cfg.DatabaseDsn)
		if err != nil {
			log.Fatal(err)
		}
		defer dbpool.Close()

		storage := storage.NewURLDBStorage(dbpool)
		err = storage.Prepare(ctx)
		if err != nil {
			log.Fatal(err)
		}
		service = app.NewURLShortener(&storage)
	case cfg.FileStoragePath != "":
		storage := storage.NewURLFileStorage(&cfg)
		storage.LoadData()
		service = app.NewURLShortener(&storage)
	default:
		storage := storage.NewURLMemStorage()
		service = app.NewURLShortener(&storage)
	}

	server := http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler.NewURLHandler(&service, &cfg, dbpool),
	}

	go func() {
		log.Println("server starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-quit

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
