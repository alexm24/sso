package main

import (
	"log/slog"
	"os"

	"github.com/alexm24/sso/internal/app"
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

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCServer.MustRun()
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
