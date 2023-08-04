package client

import (
	"context"
	"log"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/client/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/pb"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(cfg *config.Config) (interfaces.ProductClient, error) {

	gcc, err := grpc.Dial(cfg.PRODUCT_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicln("unable to connect product server!")
		return nil, err
	}
	client := pb.NewProductServiceClient(gcc)
	return &ProductClient{
		client: client,
	}, nil
}

func (p *ProductClient) AddProduct(ctx context.Context, product models.Product) error {
	return nil
}

func (p *ProductClient) AddCategory(ctx context.Context, Category request.CategoryReq) error {
	return nil
}

func (p *ProductClient) GetAllBrands(ctx context.Context) (brand []response.Brand, err error) {
	return []response.Brand{}, nil
}

func (p *ProductClient) GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {
	return []response.ResponseProduct{}, nil
}

func (p *ProductClient) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {
	return nil
}
func (p *ProductClient) GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, count int, err error) {
	return []response.ProductItemResp{}, 0, nil
}

func (p *ProductClient) UpdateProduct(ctx context.Context, product models.Product) error {
	return nil
}

func (p *ProductClient) DeleteProduct(ctx context.Context, productID uint) (models.Product, error) {
	return models.Product{}, nil
}
func (p *ProductClient) SKUhelper(ctx context.Context, productId uint) (interface{}, error) {
	return nil, nil
}
