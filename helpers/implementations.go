package helpers

import (
	"context"
	"errors"
	"go_microservice/encrypt_string"
)

var errEmpty = errors.New("Secret key ortext should not be empty")

type EncryptServiceInstance struct {}

func (EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {
	encrypted := encrypt_string.EncryptString(key, text)

	return encrypted, nil
}

func (EncryptServiceInstance) Decrypt(_ context.Context, key string, text string) (string, error) {
	if key == "" || text == "" {
		return "", errEmpty
	}

	decrypted := encrypt_string.DecryptString(key, text)

	return decrypted, nil
}

