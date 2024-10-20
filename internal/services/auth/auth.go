package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Andrew-Savin-msk/sso/internal/lib/jwt"
	"github.com/Andrew-Savin-msk/sso/internal/services"
	"github.com/Andrew-Savin-msk/sso/internal/store"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Refactor errors system

type Auth struct {
	log          *slog.Logger
	userSaver    store.UserSaver
	userProvider store.UserProvider
	appProvider  store.AppProvider
	TokenTTL     time.Duration
}

func New(log *slog.Logger, userSaver store.UserSaver, userProvider store.UserProvider, appProvider store.AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		TokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email, pass string, appId int) (string, error) {
	const op = "auth.Login"

	log := a.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("attempting to login user")

	user, err := a.userProvider.Get(ctx, email)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			a.log.Warn("user not found with error: %w", err)

			return "", fmt.Errorf("%s: %w", op, services.ErrInvalidCredentials)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswdHash, []byte(pass))
	if err != nil {
		a.log.Info("invalid credentials", err)

		return "", fmt.Errorf("%s: %w", op, services.ErrInvalidCredentials)
	}

	app, err := a.appProvider.App(ctx, appId)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.NewToken(user, app, a.TokenTTL)
	if err != nil {
		a.log.Info("failed to generate token with error: ", err)

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) Register(ctx context.Context, email, pass string) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("registering new user")

	pHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash with error: %w", err)

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.userSaver.Save(ctx, email, pHash)
	if err != nil {

		if errors.Is(err, store.ErrUserExists) {
			a.log.Warn("user not found with error: %w", err)

			return 0, fmt.Errorf("%s: %w", op, services.ErrInvalidCredentials)
		}

		log.Error("failed to save user with error: %w", err)

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return id, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	const op = "auth.Login"

	log := a.log.With(slog.String("op", op), slog.Int64("user_id", userId))

	log.Info("checking if user is admin")

	isAdmin, err := a.userProvider.IsAdmin(ctx, userId)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			a.log.Warn("user not found with error: %w", err)

			return false, fmt.Errorf("%s: %w", op, services.ErrInvalidCredentials)
		}

		if errors.Is(err, store.ErrAppNotFound) {
			a.log.Warn("app not found with error: %w", err)

			return false, fmt.Errorf("%s: %w", op, services.ErrInvalidAppId)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))

	return isAdmin, nil
}
