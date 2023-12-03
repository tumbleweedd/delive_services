package auth_server

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/auth"
	customErrors "github.com/tumbleweedd/delive_services/sso/internal/lib/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterRequestValidation struct {
	Name            string `validate:"required"`
	Lastname        string
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=8,max=20"`
	PasswordConfirm string `validate:"required,eqfield=Password"`
	OfficeUUID      string `validate:"required,uuid"`
}

func (regReqValid *RegisterRequestValidation) Validate() error {
	return validator.New().Struct(regReqValid)
}

func registerRequestValidationFromGRPSRequest(grpcRequest *auth.RegisterRequest) *RegisterRequestValidation {
	return &RegisterRequestValidation{
		Name:            grpcRequest.GetName(),
		Lastname:        grpcRequest.GetLastName(),
		Email:           grpcRequest.GetEmail(),
		Password:        grpcRequest.GetPassword(),
		PasswordConfirm: grpcRequest.GetPasswordConfirm(),
		OfficeUUID:      grpcRequest.GetOfficeUuid(),
	}
}

func (authApi *AuthServerAPI) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	validRequest := registerRequestValidationFromGRPSRequest(request)

	if err := validRequest.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	userUUID, err := authApi.authService.RegisterNewUser(
		ctx,
		validRequest.Name,
		validRequest.Lastname,
		validRequest.Email,
		validRequest.OfficeUUID,
		validRequest.Password, validRequest.PasswordConfirm,
	)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &auth.RegisterResponse{
		UserUuid: userUUID.String(),
	}, nil
}
