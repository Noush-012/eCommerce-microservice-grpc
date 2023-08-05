package service

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/pb"
	interfaces "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/usecase/interface"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/utils/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
func (a *authServiceServer) SignUp(ctx context.Context, req *pb.SignupDataRequest) (*emptypb.Empty, error) {
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
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (a *authServiceServer) Login(ctx context.Context, req *pb.LoginDataRequest) (*pb.LoginDataResponse, error) {

	resp, err := a.usecase.Login(ctx, models.Users{
		UserName: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &pb.LoginDataResponse{}, err
	}
	return &pb.LoginDataResponse{
		Id: uint32(resp.ID),
	}, nil

}

func (a *authServiceServer) OTPLogin(ctx context.Context, req *pb.OTPLoginDataRequest) (*pb.OTPLoginDataResponse, error) {
	resp, err := a.usecase.OTPLogin(&ctx, request.OTPVerify{
		OTP:    req.Otp,
		UserID: uint(req.UserID),
	})
	if err != nil {
		return &pb.OTPLoginDataResponse{}, err
	}

	return &pb.OTPLoginDataResponse{
		Id:        int32(resp.ID),
		UserName:  resp.UserName,
		FirstName: resp.FirstName,
		Email:     resp.Email,
		Phone:     resp.Phone,
	}, nil
}
