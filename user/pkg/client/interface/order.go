package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/pb"
)

type OrderClient interface {
	// GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error)
	GetOrderHistory(ctx context.Context, orderReq *pb.GetOrderHistoryRequest) (*pb.GetOrderHistoryResponse, error)
}
