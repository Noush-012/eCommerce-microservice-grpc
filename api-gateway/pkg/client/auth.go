package client

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/client/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(cfg *config.Config) (interfaces.AuthClient, error) {

	gcc, err := grpc.Dial(cfg.AUTH_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
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
