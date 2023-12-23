package auth_server_grpc

import (
	"context"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
)

type LogoutRequestValidation struct {
}

func (authApi *AuthServerAPI) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}
