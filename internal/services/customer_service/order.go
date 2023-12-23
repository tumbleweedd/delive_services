package customer_service

import (
	"context"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"log/slog"
)

type Order interface {
	CreateOrder(ctx context.Context, userUUID string, orderItems ...[]*models.OrderItem) error
	GetActualMenu(ctx context.Context) ([]*models.Product, error)
}

type OrderService struct {
	log       *slog.Logger
	orderRepo repository.Order
}

func NewOrderService(log *slog.Logger, orderRepo repository.Order) *OrderService {
	return &OrderService{
		log:       log,
		orderRepo: orderRepo,
	}
}

func (orderService *OrderService) CreateOrder(ctx context.Context, userUUID string, orderItems ...[]*models.OrderItem) error {
	const op = "services.customer.CreateOrder"

	//log := orderService.log.With(slog.String("op", op), slog.String("userGUID", userUUID))
	//
	//userUUIDParsed, err := uuid.Parse(userUUID)
	//if err != nil {
	//	log.Error("failed to parse user uuid", sl.Err(err))
	//	return fmt.Errorf("%s: %w", op, err)
	//}
	//
	//return orderService.orderRepo.CreateOrder(ctx, userUUIDParsed, orderItems...)
	return nil
}
