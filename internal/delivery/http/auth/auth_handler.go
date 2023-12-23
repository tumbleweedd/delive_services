package auth_http

import (
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/delive_services/sso/internal/services/auth_service"
	"log/slog"
)

type Initializer interface {
	InitRoutes(api *gin.RouterGroup)
}

type AuthHandlerInterface interface {
	Initializer
}

type AuthHandler struct {
	log          *slog.Logger
	authServices auth_service.Auth
}

func NewAuthHandler(log *slog.Logger, authServices auth_service.Auth) *AuthHandler {
	return &AuthHandler{
		log:          log,
		authServices: authServices,
	}
}

func (ah *AuthHandler) InitRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", ah.Login)
		auth.POST("/register", ah.Register)
		auth.POST("/refresh-token", ah.RefreshToken)
	}
}
