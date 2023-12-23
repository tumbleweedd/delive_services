package sso_grpc

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"github.com/tumbleweedd/delive_services/sso/internal/delivery/grpc/auth_server_grpc"
	"github.com/tumbleweedd/delive_services/sso/internal/delivery/grpc/customer_server"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
	"github.com/tumbleweedd/delive_services/sso/internal/services/auth_service"
	"github.com/tumbleweedd/delive_services/sso/internal/services/customer_service"
	"google.golang.org/grpc"
)

type serverAPI struct {
	authServerAPI     *auth_server_grpc.AuthServerAPI
	customerServerAPI *customer_server.CustomerServerAPI
}

func newServerAPI(
	authService auth_service.Auth,
	customerService customer_service.User,
) *serverAPI {
	return &serverAPI{
		authServerAPI:     auth_server_grpc.NewAuthServerAPI(authService),
		customerServerAPI: customer_server.NewCustomerServerAPI(customerService),
	}
}

func Register(gRPC *grpc.Server, svc *services.Services) {
	servers := newServerAPI(svc.Auth, svc.User)

	auth.RegisterAuthServer(gRPC, servers.authServerAPI)
}
