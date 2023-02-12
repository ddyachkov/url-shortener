package config

import (
	"log"

	"github.com/caarlos0/env"
)

type ServerConfig struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
}

type StorageConfig struct {
	StoragePath string `env:"FILE_STORAGE_PATH" envDefault:"./data/data.txt"`
}

func GetServerConfig() ServerConfig {
	var cfg ServerConfig

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

func GetStorageConfig() StorageConfig {
	var cfg StorageConfig

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
