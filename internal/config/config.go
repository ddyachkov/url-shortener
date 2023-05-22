package config

import (
	"flag"
	"log"
	"strings"

	"github.com/caarlos0/env"
)

var (
	serverAddress   string
	baseURL         string
	databaseDsn     string
	fileStoragePath string
	secretKey       string
	httpsEnabled    bool
)

// ServerConfig contains server configuration.
type ServerConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	SecretKey       string `env:"SECRET_KEY"`
	HTTPSEnabled    bool   `env:"ENABLE_HTTPS"`
}

// DefaultServerConfig returns ServerConfig object with values saved from env and flags.
func DefaultServerConfig() *ServerConfig {
	cfg := &ServerConfig{
		ServerAddress:   serverAddress,
		BaseURL:         baseURL,
		DatabaseDsn:     databaseDsn,
		FileStoragePath: fileStoragePath,
		SecretKey:       secretKey,
		HTTPSEnabled:    httpsEnabled,
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.HTTPSEnabled {
		cfg.BaseURL = strings.Replace(cfg.BaseURL, "http://", "https://", 1)
	}

	return cfg
}

func init() {
	flag.StringVar(&serverAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base URL")
	flag.StringVar(&databaseDsn, "d", "", "database data source name")
	flag.StringVar(&fileStoragePath, "f", "", "file storage path")
	flag.StringVar(&secretKey, "k", "thisisthirtytwobytelongsecretkey", "secret key")
	flag.BoolVar(&httpsEnabled, "s", false, "enable https")
}
