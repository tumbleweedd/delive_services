package services

import (
	auth_service "github.com/tumbleweedd/delive_services/sso/internal/services/auth"
	customer_service "github.com/tumbleweedd/delive_services/sso/internal/services/customer"
	"log/slog"
	"time"
)

type Service struct {
	log *slog.Logger

	*auth_service.AuthService
	*customer_service.CustomerService

	tokenTTL time.Duration
}

func NewService(
	log *slog.Logger,
	authRepository auth_service.Auth,
	customerRepository customer_service.CustomerRepository,
	tokenTTL time.Duration,
) *Service {
	return &Service{
		log: log,

		AuthService:     auth_service.NewAuthService(log, authRepository, tokenTTL),
		CustomerService: customer_service.NewCustomerService(log, customerRepository),

		tokenTTL: tokenTTL,
	}
}
