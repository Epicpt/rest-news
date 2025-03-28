package main

import (
	"log"

	"rest-news/config"
	"rest-news/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error creating config: %s", err)
	}

	log.Println("Config initializated")

	app.Run(cfg)
}
