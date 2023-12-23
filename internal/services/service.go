package services

import (
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"github.com/tumbleweedd/delive_services/sso/internal/services/auth_service"
	"github.com/tumbleweedd/delive_services/sso/internal/services/customer_service"
	"log/slog"
	"time"
)

type Services struct {
	log      *slog.Logger
	tokenTTL time.Duration

	auth_service.Auth

	customer_service.User
	customer_service.Order
}

func NewServices(
	log *slog.Logger,
	tokenTTL time.Duration,
	customerRepo repository.Customer,
) *Services {
	return &Services{
		log:      log,
		tokenTTL: tokenTTL,

		Auth: auth_service.NewAuthService(log, customerRepo, tokenTTL),
		User: customer_service.NewUserService(log, customerRepo),
	}
}
