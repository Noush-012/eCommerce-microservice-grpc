package client

import (
	"context"
	"fmt"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/config"
	interfaces "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/client/interface"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserServiceClient(cfg *config.Config) (interfaces.UserClient, error) {
	cc, err := grpc.Dial(cfg.USER_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	Client := pb.NewUserServiceClient(cc)
	return &UserClient{client: Client}, nil
}
func (u *UserClient) FindUser(user models.UserDataRequest) (models.User, error) {

	res, err := u.client.FindUser(context.Background(), &pb.User{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Phone:    user.Phone,
	})

	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:          int32(res.ID),
		UserName:    res.UserName,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Age:         int32(res.Age),
		Email:       res.Email,
		Phone:       res.Phone,
		BlockStatus: res.BlockStatus,
		Password:    res.Password,
	}, nil

}
