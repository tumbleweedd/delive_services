package customer_server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/tumbleweedd/delive_protos/gen/go/sso/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserListRequestValidation struct {
	OfficeUUID string `validation:"required,uuid"`
}

func (u *UserListRequestValidation) Validate() error {
	return validator.New().Struct(u)
}

func userListRequestValidationFromGRPCRequest(grpcRequest *customer.GetUserListRequest) *UserListRequestValidation {
	return &UserListRequestValidation{
		OfficeUUID: grpcRequest.GetOfficeUuid(),
	}
}

func (customerApi *CustomerServerAPI) GetUserList(ctx context.Context, request *customer.GetUserListRequest) (*customer.GetUserListResponse, error) {
	validRequest := userListRequestValidationFromGRPCRequest(request)

	if err := validRequest.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	users, err := customerApi.customerService.GetUserList(ctx, validRequest.OfficeUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	result := make([]*customer.User, 0, len(users))
	for _, user := range users {
		result = append(result, &customer.User{
			// TODO: change id to uuid
			Id:   user.UUID.String(),
			Name: user.Name,
			//Lastname: user.LastName,
			Email:      user.Email,
			OfficeName: user.OfficeName,
			OfficeUuid: user.OfficeUUID,
			CreatedAt:  timestamppb.New(user.CreatedAt),
		})
	}

	return &customer.GetUserListResponse{Data: result}, nil
}
