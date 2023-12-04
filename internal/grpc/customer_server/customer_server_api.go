package customer_server

import (
	"github.com/tumbleweedd/delive_services/sso/internal/services"
)

// TODO: разобраться с импортами сервисом из prodtos репозитория
type CustomerServerAPI struct {
	customerService services.Customer
}

func NewCustomerServerAPI(customerService services.Customer) *CustomerServerAPI {
	return &CustomerServerAPI{
		customerService: customerService,
	}
}
