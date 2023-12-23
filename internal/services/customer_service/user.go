package customer_service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
	"github.com/tumbleweedd/delive_services/sso/internal/lib/logger/sl"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"log/slog"
)

type User interface {
	GetUserList(ctx context.Context, officeUUID string) ([]*models.UserStruct, error)
}

type UserService struct {
	log                *slog.Logger
	customerRepository repository.Customer
}

func NewUserService(log *slog.Logger, customerRepository repository.Customer) *UserService {
	return &UserService{
		log:                log,
		customerRepository: customerRepository,
	}

}
func (userService *UserService) GetUserList(ctx context.Context, officeUUID string) ([]*models.UserStruct, error) {
	const op = "services.customer.GetUserList"

	log := userService.log.With(slog.String("op", op), slog.String("officeUUID", officeUUID))

	parsedOfficeUUID, err := uuid.Parse(officeUUID)
	if err != nil {
		log.Error("failed to parse office uuid", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	users, err := userService.customerRepository.GetUserList(ctx, parsedOfficeUUID)
	if err != nil {
		log.Error("failed to get user list", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}
