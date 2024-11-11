package usecase

import (
	"context"
	"errors"

	"github.com/Zainal21/renco-boilerplate/internal/domain"
	"github.com/Zainal21/renco-boilerplate/internal/dtos"
	"github.com/Zainal21/renco-boilerplate/internal/repositories"
	"github.com/Zainal21/renco-boilerplate/internal/utils"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
)

type baseAuthUsecase struct {
	env            *config.Env
	userRepository repositories.UserRepository
	jwtAuthUsecase domain.JwtAuthUseCase
}

func NewAuthUsecase(env *config.Env, userRepository repositories.UserRepository, jwtAuthUsecase domain.JwtAuthUseCase) domain.AuthUsecase {
	return &baseAuthUsecase{
		env:            env,
		userRepository: userRepository,
		jwtAuthUsecase: jwtAuthUsecase,
	}
}

func (b *baseAuthUsecase) SignUp(ctx context.Context, email, password string) error {
	user, err := b.userRepository.GetUserByEmail(email)
	if user != nil {
		return errors.New("user already exist")
	}
	if err != nil {
		return err
	}

	firebaseUID, err := b.jwtAuthUsecase.CreateUser(email, password)
	if err != nil {
		return err
	}

	metadata := utils.GenerateMetadata()
	userPayload := &dtos.CreateUserDto{
		UID:         metadata.UID(),
		FirebaseUID: firebaseUID,
		Email:       email,
		CreatedAt:   metadata.CreatedAt,
		UpdatedAt:   metadata.UpdatedAt,
	}
	_, err = b.userRepository.CreateUser(userPayload)
	if err != nil {
		return err
	}

	return nil
}

func (b *baseAuthUsecase) GetAccessToken(ctx context.Context, email, password string) (string, error) {
	user, err := b.userRepository.GetUserByEmail(email)
	if user == nil {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", err
	}

	accessToken, err := b.jwtAuthUsecase.GetAccessToken(email, password)
	if err != nil {
		return "", err
	}
	if accessToken == "" {
		return "", errors.New("invalid password")
	}

	return accessToken, nil
}
