//go:build migrate

package app

import (
	"errors"
	"flag"
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

	// Определяем флаг -direction (по умолчанию "up")
	direction := flag.String("direction", "up", "Migration direction: up or down")
	flag.Parse()

	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	databaseURL += "?sslmode=disable"

	m, err := migrate.New("file://migrations", databaseURL)
	if err != nil {
		log.Fatalf("migrate: failed to initialize: %s", err)
	}
	defer m.Close()

	switch *direction {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		log.Fatalf("Invalid migration direction: %s", *direction)
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: %s error: %s", *direction, err)
	}

	log.Printf("Migrate: %s success", *direction)
}
