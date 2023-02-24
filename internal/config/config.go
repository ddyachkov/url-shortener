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
	secretKey       string
)

type ServerConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	SecretKey       string `env:"SECRET_KEY"`
}

func NewServerConfig() ServerConfig {
	cfg := ServerConfig{
		ServerAddress:   serverAddress,
		BaseURL:         baseURL,
		FileStoragePath: fileStoragePath,
		SecretKey:       secretKey,
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

func init() {
	flag.StringVar(&serverAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base URL")
	flag.StringVar(&fileStoragePath, "f", "./data/data.txt", "file storage path")
	flag.StringVar(&secretKey, "k", "thisisthirtytwobytelongsecretkey", "secret key")
}
