package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
)

type Order interface {
	CreateOrder(ctx context.Context, order *models.Order, positions ...[]*models.OrderItem) error
	GetOrder(ctx context.Context, orderUUID uuid.UUID) (*models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) error
	DeleteOrder(ctx context.Context, orderUUID uuid.UUID) error
}

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (orderRepo *OrderRepository) GetOrder(ctx context.Context, orderUUID uuid.UUID) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (orderRepo *OrderRepository) UpdateOrder(ctx context.Context, order *models.Order) error {
	//TODO implement me
	panic("implement me")
}

func (orderRepo *OrderRepository) DeleteOrder(ctx context.Context, orderUUID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (orderRepo *OrderRepository) CreateOrder(ctx context.Context, order *models.Order, positions ...[]*models.OrderItem) error {
	//TODO implement me
	panic("implement me")
}
