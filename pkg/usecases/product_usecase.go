package usecases

import (
	"context"
	"errors"

	"github.com/landistas/enterprise-go-examples/pkg/entities"
)

var _ ProductUseCase = (*DefaultProductUseCase)(nil)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error)
	GetCatalog(ctx context.Context, filterOptions entities.CatalogFilterOptions) (*entities.Catalog, error)
	GetProduct(ctx context.Context, productID string) (*entities.Product, error)
}

type ProductStorageAdapter interface {
	Add(ctx context.Context, product entities.Product) (*entities.Product, error)
	Get(ctx context.Context, productID string) (*entities.Product, error)
	List(ctx context.Context, filterOptions entities.CatalogFilterOptions) (*entities.Catalog, error)
}

type DefaultProductUseCase struct {
	productStorageAdapter ProductStorageAdapter
}

func NewDefaultProductUseCase(productStorageAdapter ProductStorageAdapter) DefaultProductUseCase {
	return DefaultProductUseCase{productStorageAdapter: productStorageAdapter}
}

func (useCase DefaultProductUseCase) CreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {

	// * Validate the product
	// ** ID is not empty, name is not empty, price is bigger than zero
	// ** Check the product does not exist
	returnedProduct, err := useCase.productStorageAdapter.Get(ctx, product.ID)
	if err != nil {
		return nil, err
	}

	if returnedProduct != nil {
		return nil, errors.New("product already exists")
	}

	// * Store the product
	createdProduct, err := useCase.productStorageAdapter.Add(ctx, product)
	if err != nil {
		return nil, err
	}
	// Return the created product without error
	// Note: if any error return the error
	return createdProduct, nil
}

func (useCase DefaultProductUseCase) GetCatalog(ctx context.Context, filterOptions entities.CatalogFilterOptions) (*entities.Catalog, error) {
	//TODO implement
	return nil, nil
}

func (useCase DefaultProductUseCase) GetProduct(ctx context.Context, productID string) (*entities.Product, error) {
	return useCase.productStorageAdapter.Get(ctx, productID)
}
