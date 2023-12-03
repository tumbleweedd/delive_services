package repository

import (
	"github.com/jmoiron/sqlx"
	auth_repo "github.com/tumbleweedd/delive_services/sso/internal/repository/auth"
	customer_repo "github.com/tumbleweedd/delive_services/sso/internal/repository/customer"
)

type Repository struct {
	db *sqlx.DB

	*customer_repo.CustomerRepository
	*auth_repo.AuthRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:                 db,
		CustomerRepository: customer_repo.NewCustomerRepository(),
		AuthRepository:     auth_repo.NewAuthRepository(),
	}
}
