package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

var (
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
)

func init() {
	log.Println("Init config")
	cfg := struct {
		ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
		BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
		FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"./data/data.txt"`
	}{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&ServerAddress, "a", cfg.ServerAddress, "help message for flagname")
	log.Println("ServerAddress:", ServerAddress)
	flag.StringVar(&BaseURL, "b", cfg.BaseURL, "help message for flagname")
	log.Println("BaseURL:", BaseURL)
	flag.StringVar(&FileStoragePath, "f", cfg.FileStoragePath, "help message for flagname")
	log.Println("FileStoragePath:", FileStoragePath)

}
