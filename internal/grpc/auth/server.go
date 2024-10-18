package auth

import (
	"context"

	ssov1 "github.com/Andrew-Savin-msk/protos/gen/go/sso"
	"google.golang.org/grpc"
)

// THIS PACKAGE JUST IMPLEMENTS GRPC AUTH HANDLER INTERFACES
type serverApi struct {
	ssov1.UnimplementedAuthServer
}

func Register(srv *grpc.Server) {
	ssov1.RegisterAuthServer(srv, &serverApi{})
}

func (s *serverApi) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponce, error) {
	panic("unimplemented")
}

func (s *serverApi) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("unimplemented")
}

func (s *serverApi) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("unimplemented")
}
