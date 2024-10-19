package app

import (
	"log/slog"

	grpcmux "github.com/Andrew-Savin-msk/sso/internal/app/grpc_mux"
	"github.com/Andrew-Savin-msk/sso/internal/config"
)

type App struct {
	GRPCSrv *grpcmux.GrpcMux
}

func New(log *slog.Logger, cfg *config.Config) *App {
	grpcMux := grpcmux.New(log, cfg.GRPCSrv.Port)

	return &App{
		GRPCSrv: grpcMux,
	}
}
