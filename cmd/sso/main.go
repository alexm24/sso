package main

import (
	"log/slog"
	"os"

	"github.com/alexm24/sso/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// init configs
	configPath := "configs/config.yaml"
	cfg := config.MustLoad(configPath)

	log := setupLogger(cfg.Env)

	log.Info("starting", slog.String("env", cfg.Env))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
