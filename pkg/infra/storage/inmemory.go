package storage

import (
	"context"
	"errors"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/outputs/storage"
)

var _ storage.InfraStorage = (*InMemoryInfraStorage)(nil)

type InMemoryInfraStorage struct {
	kvs map[string]interface{}
}

func NewInMemoryInfraStorage() *InMemoryInfraStorage {
	return &InMemoryInfraStorage{kvs: make(map[string]interface{})}
}

func (storage *InMemoryInfraStorage) Save(ctx context.Context, id string, v interface{}) error {
	if _, ok := storage.kvs[id]; ok {
		return errors.New("already exists")
	}
	storage.kvs[id] = v
	return nil
}

func (storage InMemoryInfraStorage) Read(ctx context.Context, id string) (interface{}, error) {
	if v, ok := storage.kvs[id]; ok {
		return v, nil
	}
	return nil, nil
}

func (storage InMemoryInfraStorage) ListAll(ctx context.Context, filterFunc storage.FilterFunc) (interface{}, error) {
	vs := make([]interface{}, 0)
	for _, v := range storage.kvs {
		if !filterFunc(v) {
			vs = append(vs, v)
		}
	}
	return vs, nil
}
