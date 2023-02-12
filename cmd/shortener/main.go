package main

import (
	"flag"
	"log"

	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/server"
)

func main() {
	flag.Parse()
	log.Println("ServerAddress:", config.ServerAddress)
	log.Println("BaseURL:", config.BaseURL)
	log.Println("FileStoragePath:", config.FileStoragePath)
	server.Run()
}
