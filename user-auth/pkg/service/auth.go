package service

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/pb"
	interfaces "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServer struct {
	usecase interfaces.AuthService
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase interfaces.AuthService) pb.AuthServiceServer {
	return &authServiceServer{
		usecase: usecase,
	}
}

// user signup
func (a *authServiceServer) UserSignup(ctx context.Context, req *pb.SignupDataRequest) error {
	signupData := models.User{
		UserName:  req.GetUserName(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Age:       int32(req.Age),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassword(),
	}

	if err := a.usecase.SignUp(ctx, signupData); err != nil {
		return status.Errorf(codes.Internal, "%s", err.Error())
	}
	return nil
}
