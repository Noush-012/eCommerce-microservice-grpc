package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/repository/interfaces"
	service "github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/usecase/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/utils/response"
)

type productUseCase struct {
	ProductRepository interfaces.ProductRepository
}

func NewProductUseCase(ProdRepo interfaces.ProductRepository) service.ProductService {
	return &productUseCase{ProductRepository: ProdRepo}
}

func (p *productUseCase) AddProduct(ctx context.Context, product models.Product) error {
	// Check the product already exists in databse
	if dbProd, err := p.ProductRepository.FindProduct(ctx, product); err != nil {
		return err
	} else if dbProd.ID != 0 {
		return fmt.Errorf("product already exist with %s product name", dbProd.Name)
	}
	return p.ProductRepository.SaveProduct(ctx, product)

}
func (p *productUseCase) AddCategory(ctx context.Context, brand request.CategoryReq) error {
	// check if req brand already exists in db
	dbBrand, _ := p.ProductRepository.FindBrand(ctx, brand)
	if dbBrand.ID != 0 {
		return fmt.Errorf("brand already exist with %s name", brand.CategoryName)
	}
	if err := p.ProductRepository.AddCategory(ctx, brand); err != nil {
		return err
	}

	return nil

}

// to get all brands
func (p *productUseCase) GetAllBrands(ctx context.Context) (brand []response.Brand, err error) {
	allBrands, err := p.ProductRepository.GetAllBrand(ctx)
	if err != nil {
		return brand, err
	}
	return allBrands, nil
}

// to get all product
func (p *productUseCase) GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {
	return p.ProductRepository.GetAllProducts(ctx, page)
}

// to update product
func (p *productUseCase) UpdateProduct(ctx context.Context, product models.Product) error {
	// Check requested product is exist or not
	existingProduct, err := p.ProductRepository.FindProductByID(ctx, product.ID)
	if err != nil {
		return err
	} else if existingProduct.Name == "" {
		return errors.New("invalid product_id")
	}

	// check the given product_name already exist or not
	existingProduct, err = p.ProductRepository.FindProduct(ctx, models.Product{Name: product.Name})
	if err != nil {
		return err
	} else if existingProduct.ID != 0 && existingProduct.ID != product.ID {
		return errors.New("can't update the product \nrequested product_name already existing in database")
	}

	return p.ProductRepository.UpdateProduct(ctx, product)
}

// to delete a product
func (p *productUseCase) DeleteProduct(ctx context.Context, productID uint) (models.Product, error) {

	existingProduct, err := p.ProductRepository.DeleteProduct(ctx, productID)
	if err != nil {
		return models.Product{}, err
	}
	return existingProduct, nil
}

// to add product item for a product
func (p *productUseCase) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {
	if err := p.ProductRepository.AddProductItem(ctx, productItem); err != nil {
		return err
	}
	return nil
}

// to get a product variant
func (p *productUseCase) GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, count int, err error) {

	productItems, err := p.ProductRepository.GetProductItems(ctx, productId)
	if err != nil {
		return productItems, count, err
	}

	count = len(productItems)

	return productItems, count, nil
}

func (p *productUseCase) SKUhelper(ctx context.Context, productId uint) (interface{}, error) {
	dbProduct, err := p.ProductRepository.FindProductByID(ctx, productId)
	if err != nil {
		return "", err
	}
	return dbProduct.CategoryID, nil
}
