package auth

// THIS PACKAGE JUST IMPLEMENTS GRPC AUTH HANDLER INTERFACES

import (
	"context"

	ssov1 "github.com/Andrew-Savin-msk/protos/gen/go/sso"
	grpchandlers "github.com/Andrew-Savin-msk/sso/internal/grpc_handlers"
	"github.com/Andrew-Savin-msk/sso/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer
	auth services.Auth
}

const (
	emptyValue = 0
)

func Register(srv *grpc.Server, auth services.Auth) {
	ssov1.RegisterAuthServer(srv, &serverApi{auth: auth})
}

func (s *serverApi) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponce, error) {
	if req.GetUserId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "user_id is not provided")
	}

	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, grpchandlers.ErrInternalServiceError.Error())
	}

	return &ssov1.IsAdminResponce{
		IsAdmin: isAdmin,
	}, nil
}

func (s *serverApi) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if req.GetEmail() != "" {
		return nil, status.Error(codes.InvalidArgument, "email is not provided")
	}

	if req.GetPassword() != "" {
		return nil, status.Error(codes.InvalidArgument, "password is not provided")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app_id is not provided")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, grpchandlers.ErrInternalServiceError.Error())
	}

	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverApi) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if req.GetEmail() != "" {
		return nil, status.Error(codes.InvalidArgument, "email is not provided")
	}

	if req.GetPassword() != "" {
		return nil, status.Error(codes.InvalidArgument, "password is not provided")
	}

	userId, err := s.auth.Register(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, grpchandlers.ErrInternalServiceError.Error())
	}

	return &ssov1.RegisterResponse{
		UserId: userId,
	}, nil
}
