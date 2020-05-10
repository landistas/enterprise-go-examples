package tests

import (
	"context"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/outputs/storage"
)

type StubInfraStorage struct {
	products map[string]interface{}
}

var _ storage.InfraStorage = (*StubInfraStorage)(nil)

func (stubInfraStorage *StubInfraStorage) Save(ctx context.Context, id string, v interface{}) error {
	if stubInfraStorage.products == nil {
		stubInfraStorage.products = make(map[string]interface{})
	}
	stubInfraStorage.products[id] = v
	return nil
}
func (stubInfraStorage StubInfraStorage) Read(ctx context.Context, id string) (interface{}, error) {
	product := stubInfraStorage.products[id]
	return product, nil
}
func (stubInfraStorage StubInfraStorage) ListAll(ctx context.Context, filterFunc storage.FilterFunc) (interface{}, error) {
	return nil, nil
}
