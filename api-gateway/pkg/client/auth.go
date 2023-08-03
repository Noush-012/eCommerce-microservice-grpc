package client

import (
	"context"
	"log"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/client/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/pb"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(cfg *config.Config) (interfaces.AuthClient, error) {

	gcc, err := grpc.Dial(cfg.AUTH_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicln("unable to connect auth server!")
		return nil, err
	}
	client := pb.NewAuthServiceClient(gcc)

	return &authClient{
		client: client,
	}, nil
}

func (a *authClient) SignUp(ctx context.Context, user models.Users) error {
	_, err := a.client.SignUp(ctx, &pb.SignupDataRequest{
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       uint32(user.Age),
		Password:  user.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *authClient) Login(ctx context.Context, user models.Users) (models.Users, error) {
	resp, err := a.client.Login(ctx, &pb.LoginDataRequest{
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return models.Users{}, err
	}

	return models.Users{
		ID: uint(resp.Id),
	}, nil

}

func (a *authClient) OTPLogin(ctx context.Context, body request.OTPVerify) (models.Users, error) {

	resp, err := a.client.OTPLogin(ctx, &pb.OTPLoginDataRequest{
		Otp: body.OTP,
	})
	if err != nil {
		return models.Users{}, err
	}

	return models.Users{
		ID:        uint(resp.Id),
		UserName:  resp.UserName,
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		Email:     resp.Email,
		Phone:     resp.Phone,
	}, nil

}
