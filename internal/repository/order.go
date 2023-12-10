package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
)

type Order interface {
	CreateOrder(order *models.Order, positions ...[]*models.OrderItem) error
	GetOrder(orderUUID uuid.UUID) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(orderUUID uuid.UUID) error
}

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (orderRepo *OrderRepository) GetOrder(orderUUID uuid.UUID) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (orderRepo *OrderRepository) UpdateOrder(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}

func (orderRepo *OrderRepository) DeleteOrder(orderUUID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (orderRepo *OrderRepository) CreateOrder(order *models.Order, positions ...[]*models.OrderItem) error {
	//TODO implement me
	panic("implement me")
}
