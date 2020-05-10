package storage_test

import (
	"context"
	"testing"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/outputs/storage"
	"github.com/landistas/enterprise-go-examples/pkg/entities"
	"github.com/landistas/enterprise-go-examples/tests"
	"github.com/stretchr/testify/require"
)

func TestGetOK(t *testing.T) {
	//TODO this test fail make it work
	infraStorage := &tests.StubInfraStorage{}
	storageAdapter := storage.NewProductStorageAdapter(infraStorage)
	product := entities.Product{ID: "id"}

	err := infraStorage.Save(context.TODO(), product.ID, product)
	require.Nil(t, err)

	returnedProduct, err := storageAdapter.Get(context.TODO(), product.ID)
	require.NoError(t, err)
	require.NotNil(t, returnedProduct)
	require.Equal(t, returnedProduct.ID, product.ID)

}

func TestAddOK(t *testing.T) {
	storageAdapter := storage.NewProductStorageAdapter(&tests.StubInfraStorage{})
	product := entities.Product{ID: "id"}

	createdProduct, err := storageAdapter.Add(context.TODO(), product)

	require.NoError(t, err)
	require.NotNil(t, createdProduct)
	require.Equal(t, createdProduct.ID, product.ID)
}
