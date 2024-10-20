package grpchandlers

import "errors"

var (
	ErrInternalServiceError = errors.New("internal service error has occurred")
	ErrInvalidCredentials = errors.New("provided invalid credentials")
)
