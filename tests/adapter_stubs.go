package tests

import (
	"context"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/outputs/storage"
	"github.com/landistas/enterprise-go-examples/pkg/entities"
)

type StubProductStorageAdapter struct {
	StubInfraStorage storage.InfraStorage
}

func (stubProductStorageAdapter StubProductStorageAdapter) Add(ctx context.Context, product entities.Product) (*entities.Product, error) {
	return &product, nil
}
func (stubProductStorageAdapter StubProductStorageAdapter) Get(ctx context.Context, productID string) (*entities.Product, error) {
	v, err := stubProductStorageAdapter.StubInfraStorage.Read(ctx, productID)
	product, ok := v.(entities.Product)
	if ok {
		return &product, err
	}
	return nil, err
}
func (stubProductStorageAdapter StubProductStorageAdapter) List(ctx context.Context, filterOptions entities.CatalogFilterOptions) (*entities.Catalog, error) {
	return nil, nil
}
