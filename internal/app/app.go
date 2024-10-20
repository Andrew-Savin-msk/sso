package app

import (
	"log"
	"log/slog"

	grpcmux "github.com/Andrew-Savin-msk/sso/internal/app/grpc_mux"
	"github.com/Andrew-Savin-msk/sso/internal/config"
	"github.com/Andrew-Savin-msk/sso/internal/services/auth"
	"github.com/Andrew-Savin-msk/sso/internal/store/sqlite"
)

type App struct {
	GRPCSrv *grpcmux.GrpcMux
}

func New(logger *slog.Logger, cfg *config.Config) *App {

	st, err := sqlite.New(cfg.Db.Path)

	if err != nil {
		log.Fatal(err)
	}

	authSrv := auth.New(logger, st, st, st, cfg.App.TokenTtl)

	grpcMux := grpcmux.New(logger, authSrv, cfg.GRPCSrv.Port)

	return &App{
		GRPCSrv: grpcMux,
	}
}
