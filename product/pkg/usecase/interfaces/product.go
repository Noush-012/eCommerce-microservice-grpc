package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/utils/response"
)

type ProductService interface {
	AddProduct(ctx context.Context, product models.Product) error
	AddCategory(ctx context.Context, Category request.CategoryReq) error
	GetAllBrands(ctx context.Context) (brand []response.Brand, err error)
	GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	UpdateProduct(ctx context.Context, product models.Product) error
	DeleteProduct(ctx context.Context, productID uint) (models.Product, error)

	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, count int, err error)

	SKUhelper(ctx context.Context, productId uint) (interface{}, error)
}
