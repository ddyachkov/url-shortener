//go:build linux

package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/grpc/proto"
	gs "github.com/ddyachkov/url-shortener/internal/grpc/server"
	"github.com/ddyachkov/url-shortener/internal/handler"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	flag.Parse()
	cfg := config.DefaultServerConfig()
	log.Printf("Config: %+v\n", *cfg)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	var dbpool *pgxpool.Pool
	var urlStorage storage.URLStorage
	var err error
	switch {
	case cfg.DatabaseDsn != "":
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
		if cfg.HTTPSEnabled {
			err = server.ListenAndServeTLS("server.crt", "server.key")
		} else {
			err = server.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterShortenerServer(grpcServer, &gs.ShortenerServer{Service: service, Config: cfg})

	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()

	<-quit

	srvCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(srvCtx); err != nil {
		log.Fatal(err)
	}

	grpcServer.GracefulStop()

	log.Println("Server stopped")
}
