package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/utils/response"
)

type ProductRepository interface {
	// Product CRUD section
	GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	FindProduct(ctx context.Context, product models.Product) (models.Product, error)
	SaveProduct(ctx context.Context, product models.Product) error
	FindProductByID(ctx context.Context, productID uint) (product models.Product, err error)
	UpdateProduct(ctx context.Context, product models.Product) error
	DeleteProduct(ctx context.Context, productID uint) (models.Product, error)
	GetStockStatusByProductId(c context.Context, productId uint) (qtyLeft uint, err error)

	UpdateProductItem(ctx context.Context, UpdateData request.UpdateProductItemReq) error

	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error)

	// Brand CRUD section
	FindBrand(ctx context.Context, brand request.CategoryReq) (request.CategoryReq, error)
	AddCategory(ctx context.Context, brand request.CategoryReq) (err error)
	GetAllBrand(ctx context.Context) (brand []response.Brand, err error)
}
