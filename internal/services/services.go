package services

import "context"

type Auth interface {
	Login(ctx context.Context, email, password string, appId int) (string, error)
	Register(ctx context.Context, email, password string) (int64, error)
	IsAdmin(ctx context.Context, userId int) (bool, error)
}
