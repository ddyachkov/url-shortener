package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
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
	configFile      string
)

// ServerConfig contains server configuration.
type ServerConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS" json:"server_address,omitempty"`
	BaseURL         string `env:"BASE_URL" json:"base_url,omitempty"`
	DatabaseDsn     string `env:"DATABASE_DSN" json:"database_dsn,omitempty"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" json:"file_storage_path,omitempty"`
	SecretKey       string `env:"SECRET_KEY" json:"secret_key,omitempty"`
	HTTPSEnabled    bool   `env:"ENABLE_HTTPS" json:"enable_https,omitempty"`
}

// DefaultServerConfig returns ServerConfig object with values saved from env and flags.
func DefaultServerConfig() *ServerConfig {
	cfg := &ServerConfig{}

	if configFile != "" {
		content, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(content, &cfg)
		if err != nil {
			log.Fatal(err)
		}
	}

	if serverAddress != "" {
		cfg.ServerAddress = serverAddress
	}
	if baseURL != "" {
		cfg.BaseURL = baseURL
	}
	if databaseDsn != "" {
		cfg.DatabaseDsn = databaseDsn
	}
	if fileStoragePath != "" {
		cfg.FileStoragePath = fileStoragePath
	}
	if secretKey != "" {
		cfg.SecretKey = secretKey
	}
	if httpsEnabled {
		cfg.HTTPSEnabled = httpsEnabled
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
	flag.StringVar(&configFile, "c", "./configs/config.json", "path to configuration file")
	flag.StringVar(&databaseDsn, "d", "", "database data source name")
	flag.StringVar(&fileStoragePath, "f", "", "file storage path")
	flag.StringVar(&secretKey, "k", "thisisthirtytwobytelongsecretkey", "secret key")
	flag.BoolVar(&httpsEnabled, "s", false, "enable https")
}
