package repositories

import (
	"github.com/Zainal21/renco-boilerplate/internal/dtos"
	"github.com/Zainal21/renco-boilerplate/internal/entity"
)

type UserRepository interface {
	CreateUser(CreateUserDto *dtos.CreateUserDto) (int, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByFirebaseUID(UID string) (*entity.User, error)
	GetUserByUID(UID string) (*entity.User, error)
}
