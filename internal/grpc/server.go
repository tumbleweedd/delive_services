package authgrpc

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"github.com/tumbleweedd/delive_services/sso/internal/grpc/auth_server"
	"google.golang.org/grpc"
)

type serverAPI struct {
	authServerAPI *auth_server.AuthServerAPI
}

func newServerAPI(authService auth_server.AuthService) *serverAPI {
	return &serverAPI{
		authServerAPI: auth_server.NewAuthServerAPI(authService),
	}
}

func Register(gRPC *grpc.Server, authService auth_server.AuthService) {
	servers := newServerAPI(authService)

	auth.RegisterAuthServer(gRPC, servers.authServerAPI)
}
