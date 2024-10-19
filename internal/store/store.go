package store

import (
	"context"

	"github.com/Andrew-Savin-msk/sso/internal/domain/models"
)

type UserSaver interface {
	Save(ctx context.Context, email string, passwdHash []byte) (int64, error)
}

type UserProvider interface {
	Get(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appid int) (models.App, error)
}
