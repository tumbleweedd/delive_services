package auth_http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	AppID    int32  `json:"app_id" validate:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func toLoginResponse(accessToken, refreshToken string) *LoginResponse {
	return &LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}
}

func (lr *LoginRequest) Validate() error {
	return validator.New().Struct(lr)
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := loginRequest.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := ah.authServices.Login(ctx, loginRequest.Email, loginRequest.Password, int(loginRequest.AppID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, toLoginResponse(accessToken, refreshToken))
}
