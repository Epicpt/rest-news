package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP
	Log
	PG
}

type HTTP struct {
	Port string `env-required:"true" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" env:"LOG_LEVEL"`
}

type PG struct {
	URL     string `env-required:"true" env:"PG_URL"`
	PoolMax int    `env-required:"true" env:"PG_POOL_MAX"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Файл .env не найден")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
