package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type HashUtil interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) bool
}

type baseHashUtil struct {
}

func NewHashUtil() HashUtil {
	return &baseHashUtil{}
}

func (b *baseHashUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b *baseHashUtil) ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
