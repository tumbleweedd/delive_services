package auth_server_grpc

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"github.com/tumbleweedd/delive_services/sso/internal/services/auth_service"
)

type AuthServerAPI struct {
	auth.UnimplementedAuthServer

	authService auth_service.Auth
}

func NewAuthServerAPI(authService auth_service.Auth) *AuthServerAPI {
	return &AuthServerAPI{
		authService: authService,
	}
}
