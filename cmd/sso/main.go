package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")

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
