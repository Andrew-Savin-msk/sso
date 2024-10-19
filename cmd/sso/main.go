package main

import (
	"log/slog"
	"os"

	"github.com/Andrew-Savin-msk/sso/internal/config"
)

const (
	levelLocal = "local"
	levelDev   = "dev"
	levelProd  = "prod"
)

func main() {
	cfg := config.Load()

	log := setupLogger(cfg.App.LogLevel)

	log.Info("starting application",
		slog.String("level", cfg.App.LogLevel),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPCSrv.Port),
	)
}

func setupLogger(level string) *slog.Logger {
	var log *slog.Logger

	switch level {
	case levelLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case levelDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case levelProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
