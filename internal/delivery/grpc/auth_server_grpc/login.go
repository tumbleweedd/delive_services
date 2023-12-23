package auth_server_grpc

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginRequestValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=20"`
	AppID    int32  `validate:"required"`
}

func (loginReqValid *LoginRequestValidation) Validate() error {
	return validator.New().Struct(loginReqValid)
}

func loginRequestValidationFromGRPSRequest(grpcRequest *auth.LoginRequest) *LoginRequestValidation {
	return &LoginRequestValidation{
		Email:    grpcRequest.GetEmail(),
		Password: grpcRequest.GetPassword(),
		AppID:    grpcRequest.GetAppId(),
	}
}

func (authApi *AuthServerAPI) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	validationRequest := loginRequestValidationFromGRPSRequest(request)

	if err := validationRequest.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	accessToken, refreshToken, err := authApi.authService.Login(ctx, validationRequest.Email, validationRequest.Password, int(validationRequest.AppID))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		//RefreshToken: refreshToken,
	}, nil
}
