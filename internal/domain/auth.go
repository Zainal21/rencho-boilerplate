package domain

import (
	"context"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/auth_usecase_mock.go --fake-name AuthUsecaseMock . AuthUsecase

// Controller

// Usecase
type AuthUsecase interface {
	SignUp(ctx context.Context, email, password string) error
	GetAccessToken(ctx context.Context, email, password string) (string, error)
}
