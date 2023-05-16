//go:build linux

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

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	flag.Parse()
	cfg := config.DefaultServerConfig()
	log.Printf("Config: %+v\n", *cfg)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	var dbpool *pgxpool.Pool
	var urlStorage storage.URLStorage
	switch {
	case cfg.DatabaseDsn != "":
		var err error

		dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		dbpool, err = pgxpool.New(dbCtx, cfg.DatabaseDsn)
		if err != nil {
			log.Fatal(err)
		}
		defer dbpool.Close()

		urlStorage, err = storage.NewURLDBStorage(dbpool, dbCtx)
		if err != nil {
			log.Fatalln(err.Error())
		}
	case cfg.FileStoragePath != "":
		urlStorage = storage.NewURLFileStorage(cfg)
	default:
		urlStorage = storage.NewURLMemStorage()

	}
	service := app.NewURLShortener(urlStorage)

	server := http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler.NewURLHandler(service, cfg, dbpool),
	}

	go func() {
		log.Println("Server starting...")
		log.Println("Build version:", buildVersion)
		log.Println("Build date:", buildDate)
		log.Println("Build commit:", buildCommit)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-quit

	srvCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(srvCtx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped")
}
