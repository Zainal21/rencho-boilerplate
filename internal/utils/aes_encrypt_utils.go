package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type AesEncryptUtil interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
}

type baseAesEncrypt struct {
	key []byte
}

func NewAesEncrypt(_key string) AesEncryptUtil {
	key, _ := hex.DecodeString(_key)

	return &baseAesEncrypt{
		key: key,
	}
}

func (b *baseAesEncrypt) Encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	byt := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(byt), nil
}

func (b *baseAesEncrypt) Decrypt(ciphertextStr string) (string, error) {
	ciphertextByt, err := hex.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertextByt) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertextByt[:nonceSize], ciphertextByt[nonceSize:]

	byt, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}
