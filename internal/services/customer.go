package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
	"github.com/tumbleweedd/delive_services/sso/internal/lib/logger/sl"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"log/slog"
)

type Customer interface {
	GetUserList(ctx context.Context, officeUUID string) ([]*models.UserStruct, error)
}

type CustomerService struct {
	log                *slog.Logger
	customerRepository repository.Customer
}

func (customerService *CustomerService) GetUserList(ctx context.Context, officeUUID string) ([]*models.UserStruct, error) {
	const op = "services.customer.GetUserList"

	log := customerService.log.With(slog.String("op", op), slog.String("officeUUID", officeUUID))

	parsedOfficeUUID, err := uuid.Parse(officeUUID)
	if err != nil {
		log.Error("failed to parse office uuid", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	users, err := customerService.customerRepository.GetUserList(ctx, parsedOfficeUUID)
	if err != nil {
		log.Error("failed to get user list", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

func NewCustomerService(log *slog.Logger, customerRepository repository.Customer) *CustomerService {
	return &CustomerService{
		log:                log,
		customerRepository: customerRepository,
	}
}
