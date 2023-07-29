package client

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/client/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/pb"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	client pb.UserServiceClient
}

func NewUserClient(cfg *config.Config) (interfaces.UserClient, error) {
	gcc, err := grpc.Dial(cfg.USER_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(gcc)

	return &userClient{
		client: client,
	}, nil
}

func (u *userClient) Profile(ctx context.Context, userId uint) (response.Profile, error) {
	res, err := u.client.Profile(ctx, &pb.UserId{
		UserId: uint32(userId),
	})
	if err != nil {
		return response.Profile{}, nil
	}
	profile := response.Profile{
		ID:        uint(res.ID),
		UserName:  res.UserName,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Age:       uint(res.Age),
		Email:     res.Email,
		Phone:     res.Phone,
		DefaultAddress: response.Address{
			ID:           uint(res.DefaultAddress.ID),
			UserID:       uint(res.DefaultAddress.UserID),
			House:        res.DefaultAddress.House,
			AddressLine1: res.DefaultAddress.AddressLine1,
			AddressLine2: res.DefaultAddress.AddressLine2,
			City:         res.DefaultAddress.City,
			State:        res.DefaultAddress.State,
			Country:      res.DefaultAddress.Country,
			ZipCode:      res.DefaultAddress.ZipCode,
		},
	}

	for i, v := range res.OrderHistoryList {

		profile.OrderHistory[i].ID = uint(v.ID)
		profile.OrderHistory[i].OrderDate = v.OrderDate.AsTime()
		profile.OrderHistory[i].OrderStatus = v.OrderStatus
		profile.OrderHistory[i].DeliveryStatus = v.DeliveryStatus
		profile.OrderHistory[i].OrderTotal = float64(v.OrderTotal)
		profile.OrderHistory[i].PaymentMethod = v.PaymentMethod
		profile.OrderHistory[i].PaymentStatus = v.PaymentStatus
	}

	return profile, nil

}
