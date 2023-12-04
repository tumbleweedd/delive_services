package authgrpc

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"github.com/tumbleweedd/delive_services/sso/internal/grpc/auth_server"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
	"google.golang.org/grpc"
)

type serverAPI struct {
	authServerAPI *auth_server.AuthServerAPI
}

func newServerAPI(
	authService services.Auth,
	customerService services.Customer,
) *serverAPI {
	return &serverAPI{
		authServerAPI: auth_server.NewAuthServerAPI(authService),
	}
}

func Register(gRPC *grpc.Server, svc *services.Service) {
	servers := newServerAPI(svc.Auth)

	auth.RegisterAuthServer(gRPC, servers.authServerAPI)
}
