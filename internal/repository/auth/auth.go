package auth

import (
	"context"
	"github.com/google/uuid"
)

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (authRepo *AuthRepository) SaveUser(ctx context.Context, name, lastname, email, officeUUID, pwd, pwdConfirm string) (userUUID uuid.UUID, err error) {
	return [16]byte{}, err
}
