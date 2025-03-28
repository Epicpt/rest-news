package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func New(level string) zerolog.Logger {
	// Парсим уровень логирования
	logLevel, err := zerolog.ParseLevel(strings.ToLower(level))
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	// Настраиваем логгер с выводом в stdout
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Level(logLevel)

	// Устанавливаем человекочитаемый вывод, если не продакшн
	if logLevel <= zerolog.DebugLevel {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}

	return logger
}
