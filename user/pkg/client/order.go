package client

import (
	"context"

	interfaces "github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/client/interface"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	client pb.OrderServiceClient
}

func NewOrderClient(cfg *config.Config) (interfaces.OrderClient, error) {

	gcc, err := grpc.Dial(cfg.ORDER_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewOrderServiceClient(gcc)

	return &OrderClient{
		client: client,
	}, nil
}

// func (o *OrderClient) GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error) {
// 	res, err := o.client.GetOrderHistory(ctx, &pb.GetOrderHistoryRequest{
// 		UserId:     uint32(userId),
// 		Pagination: &pb.Pagination{PageNumber: uint32(page.Count), Count: uint32(page.Count)},
// 	})
// 	if err != nil {
// 		return []response.OrderHistory{}, err
// 	}
// 	for i := 0; i < len(res.OrderHistoryList); i++ {
// 		orderHisory[i].ID = uint32(res.OrderHistoryList[i].ID)
// 		orderHisory[i].OrderDate = time.Unix(res.OrderHistoryList[i].OrderDate.GetSeconds(), int64(res.OrderHistoryList[i].OrderDate.GetNanos()))
// 		orderHisory[i].OrderStatus = res.OrderHistoryList[i].OrderStatus
// 		orderHisory[i].DeliveryStatus = res.OrderHistoryList[i].DeliveryStatus
// 		orderHisory[i].OrderTotal = float64(res.OrderHistoryList[i].OrderTotal)
// 		orderHisory[i].PaymentMethod = res.OrderHistoryList[i].PaymentMethod
// 		orderHisory[i].PaymentStatus = res.OrderHistoryList[i].PaymentStatus
// 	}

//		return orderHisory, nil
//	}
func (o *OrderClient) GetOrderHistory(ctx context.Context, orderReq *pb.GetOrderHistoryRequest) (*pb.GetOrderHistoryResponse, error) {
	return &pb.GetOrderHistoryResponse{}, nil
}
