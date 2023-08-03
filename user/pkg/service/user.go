package service

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/pb"
	userRepo "github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/repository/interface"
	usecase "github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	usecase  usecase.UserService
	userRepo userRepo.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(usecase usecase.UserService, userRepo userRepo.UserRepository) pb.UserServiceServer {
	return &UserServiceServer{
		usecase:  usecase,
		userRepo: userRepo,
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
func (u *UserServiceServer) FindUser(c context.Context, user *pb.User) (*pb.User, error) {
	dbUser, err := u.userRepo.FindUser(c, models.User{
		ID:          int32(user.ID),
		UserName:    user.UserName,
		Email:       user.Email,
		Phone:       user.Phone,
		BlockStatus: user.BlockStatus,
	})
	if err != nil {
		return &pb.User{}, err
	}

	return &pb.User{
		ID:          uint32(dbUser.ID),
		UserName:    dbUser.UserName,
		FirstName:   dbUser.FirstName,
		LastName:    dbUser.LastName,
		Age:         uint32(dbUser.Age),
		Email:       dbUser.Email,
		Phone:       dbUser.Phone,
		BlockStatus: dbUser.BlockStatus,
		Password:    dbUser.Password,
	}, nil
}
