package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

var (
	serverAddress   string
	baseURL         string
	databaseDsn     string
	fileStoragePath string
	secretKey       string
)

type ServerConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	SecretKey       string `env:"SECRET_KEY"`
}

func DefaultServerConfig() *ServerConfig {
	cfg := &ServerConfig{
		ServerAddress:   serverAddress,
		BaseURL:         baseURL,
		DatabaseDsn:     databaseDsn,
		FileStoragePath: fileStoragePath,
		SecretKey:       secretKey,
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

func init() {
	flag.StringVar(&serverAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base URL")
	flag.StringVar(&databaseDsn, "d", "", "database data source name")
	flag.StringVar(&fileStoragePath, "f", "", "file storage path")
	flag.StringVar(&secretKey, "k", "thisisthirtytwobytelongsecretkey", "secret key")
}
