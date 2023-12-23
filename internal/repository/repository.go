package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Customer
	Office
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Customer: NewUserRepository(db),
		Office:   NewOfficeRepository(db),
		Order:    NewOrderRepository(db),
	}
}
