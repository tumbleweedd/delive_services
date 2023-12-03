package customer

import "log/slog"

type CustomerRepository interface {
}

type CustomerService struct {
	log                *slog.Logger
	customerRepository CustomerRepository
}

func NewCustomerService(log *slog.Logger, customerRepository CustomerRepository) *CustomerService {
	return &CustomerService{
		log:                log,
		customerRepository: customerRepository,
	}
}
