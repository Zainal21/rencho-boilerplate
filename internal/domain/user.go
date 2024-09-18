package domain

import (
	"context"

	"github.com/Zainal21/renco-boilerplate/internal/entity"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/user_usecase_mock.go --fake-name UserUsecaseMock . UserUsecase
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/user_repository_mock.go --fake-name UserRepositoryMock . UserRepository

// Usecase
type UserUsecase interface {
	GetUserByFirebaseUID(ctx context.Context, UID string) (*entity.User, error)
	GetUserByUID(ctx context.Context, UID string) (*entity.User, error)
}
