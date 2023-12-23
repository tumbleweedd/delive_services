package customer_server

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/customer"
	"github.com/tumbleweedd/delive_services/sso/internal/services/customer_service"
)

type CustomerServerAPI struct {
	customer.UnimplementedUserServiceServer
	customer.UnimplementedOfficeServiceServer
	customer.UnimplementedOrderServiceServer

	customerService customer_service.User
}

func NewCustomerServerAPI(customerService customer_service.User) *CustomerServerAPI {
	return &CustomerServerAPI{
		customerService: customerService,
	}
}
