package service

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/pb"
	usecase "github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	usecase usecase.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(usecase usecase.UserService) pb.UserServiceServer {
	return &UserServiceServer{
		usecase: usecase,
	}
}

func (u *UserServiceServer) Profile(ctx context.Context, userId *pb.UserId) (*pb.ProfileResponse, error) {
	profileRes, err := u.usecase.Profile(ctx, uint(userId.UserId))

	if err != nil {
		errorCode := codes.Internal
		return &pb.ProfileResponse{}, status.Error(errorCode, err.Error())
	}

	return &pb.ProfileResponse{
		ID:        uint32(profileRes.ID),
		UserName:  profileRes.UserName,
		FirstName: profileRes.FirstName,
		LastName:  profileRes.LastName,
		Age:       uint32(profileRes.Age),
		Email:     profileRes.Email,
		Phone:     profileRes.Phone,
		DefaultAddress: &pb.Address{
			ID:           uint32(profileRes.DefaultAddress.ID),
			UserID:       uint32(profileRes.DefaultAddress.UserID),
			House:        profileRes.DefaultAddress.House,
			AddressLine1: profileRes.DefaultAddress.AddressLine1,
			AddressLine2: profileRes.DefaultAddress.AddressLine2,
			City:         profileRes.DefaultAddress.City,
			State:        profileRes.DefaultAddress.State,
			ZipCode:      profileRes.DefaultAddress.ZipCode,
			Country:      profileRes.DefaultAddress.Country,
		},
	}, nil

}
