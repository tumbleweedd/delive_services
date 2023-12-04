package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	auth_service "github.com/tumbleweedd/delive_services/sso/internal/services/auth"
	customer_service "github.com/tumbleweedd/delive_services/sso/internal/services/customer"
	"log/slog"
	"time"
)

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int32) (token string, err error)
	RegisterNewUser(ctx context.Context, name string, lastname string, email string, officeUUID string, pwd string, pwdConfirm string) (userUUID uuid.UUID, err error)
	IsAdmin(ctx context.Context, userGUID string) (bool, error)
	Logout(ctx context.Context, token string) (success bool) // TODO: непонятно пока, как реализовывать
	RefreshToken(ctx context.Context, userUUID string, refreshToken string) (aToken string, rToken string, err error)
}

type Customer interface {
}

type Service struct {
	log      *slog.Logger
	tokenTTL time.Duration

	Auth
	Customer
}

func NewService(
	log *slog.Logger,
	tokenTTL time.Duration,
	authRepo repository.Auth,
	customerRepo repository.Customer,
) *Service {
	return &Service{
		log:      log,
		tokenTTL: tokenTTL,

		Auth:     auth_service.NewAuthService(log, authRepo, tokenTTL),
		Customer: customer_service.NewCustomerService(log, customerRepo),
	}
}
