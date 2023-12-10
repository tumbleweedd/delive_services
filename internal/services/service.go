package services

import (
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"log/slog"
	"time"
)

type Service struct {
	log      *slog.Logger
	tokenTTL time.Duration

	Auth
	Customer
}

func NewService(
	log *slog.Logger,
	tokenTTL time.Duration,
	customerRepo repository.Customer,
) *Service {
	return &Service{
		log:      log,
		tokenTTL: tokenTTL,

		Auth:     NewAuthService(log, customerRepo, tokenTTL),
		Customer: NewCustomerService(log, customerRepo),
	}
}
