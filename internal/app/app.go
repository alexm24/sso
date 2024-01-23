package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/alexm24/sso/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	GRPCServer := grpcapp.New(log, grpcPort)

	return &App{GRPCServer}
}
