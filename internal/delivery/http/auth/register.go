package auth_http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	OfficeUUID      string `json:"office_uuid" validate:"required,uuid"`
}

func (rr *RegisterRequest) Validate() error {
	return validator.New().Struct(rr)
}

type RegisterResponse struct {
	UserUuid string `json:"user_uuid"`
}

func toRegisterResponse(userUUID string) RegisterResponse {
	return RegisterResponse{
		UserUuid: userUUID,
	}
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
	var registerRequest RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := registerRequest.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUUID, err := ah.authServices.RegisterNewUser(
		ctx,
		registerRequest.Name,
		registerRequest.Lastname,
		registerRequest.Email,
		registerRequest.OfficeUUID,
		registerRequest.Password,
		registerRequest.PasswordConfirm,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := toRegisterResponse(userUUID.String())
	ctx.JSON(http.StatusOK, response)
}
