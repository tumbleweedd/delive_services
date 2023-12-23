package sso_http

import (
	"github.com/gin-gonic/gin"
	auth_http "github.com/tumbleweedd/delive_services/sso/internal/delivery/http/auth"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
	"log/slog"
)

type Handler struct {
	log      *slog.Logger
	services *services.Services

	authHandler auth_http.AuthHandlerInterface
}

func NewHandler(log *slog.Logger, services *services.Services) *Handler {
	return &Handler{
		log:      log,
		services: services,

		authHandler: auth_http.NewAuthHandler(log, services.Auth),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	customer := router.Group("/customer")
	{
		h.authHandler.InitRoutes(customer)

		//users := customer.Group("/users")
		//{
		//	h.userHandler.InitRoutes(users)
		//}
		//
		//offices := customer.Group("/offices")
		//{
		//	h.officeHandler.InitRoutes(offices)
		//}
		//
		//orders := customer.Group("/orders")
		//{
		//	h.orderHandler.InitRoutes(orders)
		//}
	}

	return router
}
