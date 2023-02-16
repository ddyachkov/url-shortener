package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

var (
	serverAddress   string
	baseURL         string
	fileStoragePath string
)

type ServerConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

func NewServerConfig() ServerConfig {
	cfg := ServerConfig{
		ServerAddress:   serverAddress,
		BaseURL:         baseURL,
		FileStoragePath: fileStoragePath,
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

func init() {
	flag.StringVar(&serverAddress, "a", "localhost:8080", "help message for flagname")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "help message for flagname")
	flag.StringVar(&fileStoragePath, "f", "./data/data.txt", "help message for flagname")
}
