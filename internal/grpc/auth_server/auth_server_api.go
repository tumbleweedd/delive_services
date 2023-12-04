package auth_server

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
)

type AuthServerAPI struct {
	auth.UnimplementedAuthServer

	authService services.Auth
}

func NewAuthServerAPI(authService services.Auth) *AuthServerAPI {
	return &AuthServerAPI{
		authService: authService,
	}
}
