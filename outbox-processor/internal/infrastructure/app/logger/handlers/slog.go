package handlers

import (
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app"
	"github.com/D1sordxr/simple-bank/outbox-processor/internal/infrastructure/app/logger/handlers/designed"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func NewSlogLogger(config *app.Config) *slog.Logger {
	var logger *slog.Logger
	var handler slog.Handler

	switch config.AppConfig.Mode {
	case envLocal:
		logger = designed.NewPrettySlog()
		return logger
	case envDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	logger = slog.New(handler)

	return logger
}
