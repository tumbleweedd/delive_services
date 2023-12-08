package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
)

type Customer interface {
	GetUser(ctx context.Context, email string) (*models.UserStruct, error)
	SaveUser(ctx context.Context, officeUUID uuid.UUID, name, lastname, email, pwd string) (userUUID uuid.UUID, err error)
	IsAdmin(ctx context.Context, userUUID uuid.UUID) (bool, error)
	AppInfo(ctx context.Context, appID int) (*models.App, error)
}

type Office interface {
	GetOffice(officeUUID uuid.UUID) (*models.Office, error)
	CreateOffice(uuid uuid.UUID, office *models.Office) error
	GetOfficeList() ([]*models.Office, error)
	UpdateOffice(office *models.Office) error
	DeleteOffice(officeUUID uuid.UUID) error
}

type Order interface {
	CreateOrder(order *models.Order, positions ...[]*models.OrderItem) error
	GetOrder(orderUUID uuid.UUID) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(orderUUID uuid.UUID) error
}

type Repository struct {
	Customer
	Office
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Customer: NewCustomerRepository(db),
		Office:   NewOfficeRepository(db),
		Order:    NewOrderRepository(db),
	}
}
