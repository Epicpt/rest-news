//go:build migrate

package app

import (
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}

	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	databaseURL += "?sslmode=disable"

	m, err := migrate.New("file://migrations", databaseURL)
	if err != nil {
		log.Fatalf("migrate: failed to initialize: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
