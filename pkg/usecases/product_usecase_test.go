package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/landistas/enterprise-go-examples/pkg/entities"
	"github.com/landistas/enterprise-go-examples/pkg/usecases"
	"github.com/landistas/enterprise-go-examples/tests"
)

func TestCreateProductOK(t *testing.T) {
	product := entities.Product{ID: "id"}
	useCase := usecases.NewDefaultProductUseCase(
		tests.StubProductStorageAdapter{
			StubInfraStorage: &tests.StubInfraStorage{},
		},
	)

	returnedProduct, err := useCase.CreateProduct(context.TODO(), product)

	require.NoError(t, err)
	require.NotNil(t, returnedProduct)
	require.Equal(t, product.ID, returnedProduct.ID)
}

func TestCreateDuplicatedProduct(t *testing.T) {
	product := entities.Product{ID: "id"}
	infraStorage := &tests.StubInfraStorage{}
	useCase := usecases.NewDefaultProductUseCase(
		tests.StubProductStorageAdapter{
			StubInfraStorage: infraStorage,
		},
	)
	require.NoError(t, infraStorage.Save(context.TODO(), product.ID, product))

	returnedProduct, err := useCase.CreateProduct(context.TODO(), product)

	require.Nil(t, returnedProduct)
	require.Error(t, err)
}
