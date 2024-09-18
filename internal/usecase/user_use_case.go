package usecase

import (
	"context"

	"github.com/Zainal21/renco-boilerplate/internal/domain"
	"github.com/Zainal21/renco-boilerplate/internal/entity"
	"github.com/Zainal21/renco-boilerplate/internal/repositories"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
)

type baseUserUsecase struct {
	env            *config.Env
	userRepository repositories.UserRepository
}

func NewUserUsecase(env *config.Env, userRepository repositories.UserRepository) domain.UserUsecase {
	return &baseUserUsecase{
		env:            env,
		userRepository: userRepository,
	}
}

func (b *baseUserUsecase) GetUserByFirebaseUID(ctx context.Context, UID string) (*entity.User, error) {
	user, err := b.userRepository.GetUserByFirebaseUID(UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (b *baseUserUsecase) GetUserByUID(ctx context.Context, UID string) (*entity.User, error) {
	user, err := b.userRepository.GetUserByUID(UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
