package auth_server

import (
	"context"
	"github.com/google/uuid"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string, appID int32) (token string, err error)
	RegisterNewUser(ctx context.Context, name string, lastname string, email string, officeUUID string, pwd string, pwdConfirm string) (userUUID uuid.UUID, err error)
	IsAdmin(ctx context.Context, userGUID string) (bool, error)
	Logout(ctx context.Context, token string) (success bool) // TODO: непонятно пока, как реализовывать
	RefreshToken(ctx context.Context, userUUID string, refreshToken string) (aToken string, rToken string, err error)
}

type AuthServerAPI struct {
	auth.UnimplementedAuthServer

	authService AuthService
}

func NewAuthServerAPI(authService AuthService) *AuthServerAPI {
	return &AuthServerAPI{
		authService: authService,
	}
}
