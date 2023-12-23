package auth_http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RefreshTokenRequest struct {
	UserUUID     string `json:"user_uuid" validation:"required,uuid"`
	RefreshToken string `json:"refresh_token" validation:"required"`
}

func (rtr *RefreshTokenRequest) Validate() error {
	return validator.New().Struct(rtr)
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func toRefreshTokenResponse(accessToken, refreshToken string) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	var refreshTokenRequest RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&refreshTokenRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := refreshTokenRequest.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := ah.authServices.RefreshToken(
		ctx,
		refreshTokenRequest.UserUUID,
		refreshTokenRequest.RefreshToken,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, toRefreshTokenResponse(accessToken, refreshToken))
}
