package auth_server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IsAdminRequestValidation struct {
	UserUUID string `validation:"required,uuid"`
}

func (isAdminReqValidation *IsAdminRequestValidation) Validate() error {
	return validator.New().Struct(isAdminReqValidation)
}

func isAdminRequestValidationFromGRPSRequest(grpcRequest *auth.IsAdminRequest) *IsAdminRequestValidation {
	return &IsAdminRequestValidation{
		UserUUID: grpcRequest.UserUuid,
	}
}

func (authApi *AuthServerAPI) IsAdmin(ctx context.Context, request *auth.IsAdminRequest) (*auth.IsAdminResponse, error) {
	validRequest := isAdminRequestValidationFromGRPSRequest(request)

	isAdmin, err := authApi.authService.IsAdmin(ctx, validRequest.UserUUID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	return &auth.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
