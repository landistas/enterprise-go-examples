package productshttp_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/inputs/productshttp"
	"github.com/landistas/enterprise-go-examples/pkg/entities"
	"github.com/landistas/enterprise-go-examples/pkg/usecases"
)

func TestDefaultProductsHttpAdapter_CreateProduct(t *testing.T) {
	uc := &productUseCaseStub{}
	adapter := productshttp.NewDefaultProductsHttpAdapter(uc)

	// prepare the response recorder
	responseRecorder := httptest.NewRecorder()

	// Prepare the request
	productToCreate := entities.Product{
		ID:           fmt.Sprintf("%v", time.Now()), //Better to use uuid generator
		Name:         "test name",
		PriceInCents: 1,
		Inventory:    5,
	}
	buffer, err := buildJsonBodyBuffer(productToCreate)
	require.NoError(t, err)
	request := httptest.NewRequest(http.MethodPost, "/products", buffer)

	adapter.CreateProduct(responseRecorder, request)

	require.Equal(t, 200, responseRecorder.Code)
	responseProduct := entities.Product{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseProduct)
	require.NoError(t, err)

	require.Equal(t, productToCreate, responseProduct)
}

func buildJsonBodyBuffer(data interface{}) (*bytes.Buffer, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	_, err = buffer.Write(dataBytes)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

var _ usecases.ProductUseCase = (*productUseCaseStub)(nil)

type productUseCaseStub struct {
}

func (p productUseCaseStub) CreateProduct(_ context.Context, product entities.Product) (*entities.Product, error) {
	return &product, nil
}

func (p productUseCaseStub) GetCatalog(_ context.Context, _ entities.CatalogFilterOptions) (*entities.Catalog, error) {
	panic("implement me")
}

func (p productUseCaseStub) GetProduct(_ context.Context, _ string) (*entities.Product, error) {
	panic("implement me")
}
