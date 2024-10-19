package services

import "context"

type Auth interface {
	Login(ctx context.Context, email, pass string, appId int) (string, error)
	Register(ctx context.Context, email, pass string) (int64, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}
