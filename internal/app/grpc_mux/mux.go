package grpcmux

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	authrpc "github.com/Andrew-Savin-msk/sso/internal/grpc_handlers/auth"
	"github.com/Andrew-Savin-msk/sso/internal/services"
	"google.golang.org/grpc"
)

type GrpcMux struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, authSrv services.Auth, port int) *GrpcMux {
	gRPCServer := grpc.NewServer()

	authrpc.Register(gRPCServer, authSrv)

	return &GrpcMux{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (m *GrpcMux) MustRun() {
	err := m.Run()
	if err != nil {
		log.Fatalf("unable to start gRPC server, ended with error: %w", err)
	}
}

func (m *GrpcMux) Run() error {
	const op = "grpcmux.Run"

	log := m.log.With(slog.String("op", op), slog.Int("port", m.port))

	log.Info("starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", m.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	err = m.gRPCServer.Serve(l)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (m *GrpcMux) Stop() {
	const op = "grpcmux.Run"

	m.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", m.port))

	m.gRPCServer.GracefulStop()
}
