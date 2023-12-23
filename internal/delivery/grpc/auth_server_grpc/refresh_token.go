package auth_server_grpc

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshTokenRequestValidation struct {
	UserUUID     string `validation:"required,uuid"`
	RefreshToken string `validation:"required"`
}

func (rtReqValid *RefreshTokenRequestValidation) Validate() error {
	return validator.New().Struct(rtReqValid)
}

func refreshTokenRequestValidationFromGRPCRequest(grpcRequest *auth.RefreshTokenRequest) *RefreshTokenRequestValidation {
	return &RefreshTokenRequestValidation{
		UserUUID:     grpcRequest.GetUuid(),
		RefreshToken: grpcRequest.GetRefreshToken(),
	}
}

func (authApi *AuthServerAPI) RefreshToken(ctx context.Context, request *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	validRequest := refreshTokenRequestValidationFromGRPCRequest(request)

	if err := validRequest.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	accessToken, refreshToken, err := authApi.authService.RefreshToken(ctx, validRequest.UserUUID, validRequest.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &auth.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
