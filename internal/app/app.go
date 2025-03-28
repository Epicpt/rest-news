package app

import (
	"os"
	"os/signal"
	"rest-news/config"
	"rest-news/pkg/httpserver"
	"rest-news/pkg/logger"
	"rest-news/pkg/postgres"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	l.Info().Msg("Logger initialized")

	pg, err := postgres.New(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		l.Fatal().Err(err).Msg("Failed to connect to PostgreSQL")
	}
	defer pg.Close()

	l.Info().Msg("PostgreSQL initialized")

	// Use case
	// ...

	// Run server
	httpServer := httpserver.New(cfg.Port)

	// Routes
	httpServer.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msgf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		l.Error().Err(err).Msg("app - Run - httpServer.Notify")
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error().Err(err).Msg("app - Run - httpServer.Shutdown")
	}
}
