package customer_server

import (
	"github.com/tumbleweedd/delive_protos/gen/go/sso/customer"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
)

type CustomerServerAPI struct {
	customer.UnimplementedUserServiceServer
	customer.UnimplementedOfficeServiceServer
	customer.UnimplementedOrderServiceServer

	customerService services.Customer
}

func NewCustomerServerAPI(customerService services.Customer) *CustomerServerAPI {
	return &CustomerServerAPI{
		customerService: customerService,
	}
}
