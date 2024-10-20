package services

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppId = errors.New("invalid app_id")
)
