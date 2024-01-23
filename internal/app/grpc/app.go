package grpc

import (
	"log/slog"

	"google.golang.org/grpc"

	authgrpc "github.com/alexm24/sso/internal/grpc/auth"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)
	return &App{
		log,
		gRPCServer,
		port,
	}
}
