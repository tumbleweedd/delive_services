package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
	auth_repo "github.com/tumbleweedd/delive_services/sso/internal/repository/auth"
	customer_repo "github.com/tumbleweedd/delive_services/sso/internal/repository/customer"
)

type Auth interface {
	SaveUser(ctx context.Context, name, lastname, email, officeUUID, pwd, pwdConfirm string) (userUUID uuid.UUID, err error)
}

type Customer interface {
	GetUser(userUUID uuid.UUID) *models.UserStruct
}

type Repository struct {
	db *sqlx.DB

	Customer
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:       db,
		Customer: customer_repo.NewCustomerRepository(),
		Auth:     auth_repo.NewAuthRepository(),
	}
}
