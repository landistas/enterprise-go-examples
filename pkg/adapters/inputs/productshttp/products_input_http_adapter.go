package productshttp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/landistas/enterprise-go-examples/pkg/entities"

	"github.com/landistas/enterprise-go-examples/pkg/usecases"
)

type ProductsHttpAdapter interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	ListProducts(writer http.ResponseWriter, request *http.Request)
}

type DefaultProductsHttpAdapter struct {
	useCase usecases.ProductUseCase
}

func NewDefaultProductsHttpAdapter(useCase usecases.ProductUseCase) *DefaultProductsHttpAdapter {
	return &DefaultProductsHttpAdapter{useCase: useCase}
}

func (adapter *DefaultProductsHttpAdapter) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	if err != nil {
		handleErrorResponse(writer, err)
		return
	}

	createdProduct, err := adapter.useCase.CreateProduct(ctx, product)
	if err != nil {
		handleErrorResponse(writer, err)
		return
	}

	handleResponse(writer, createdProduct, nil)
}

// ListProducts lists products using the getcatalog usecase functionality
func (adapter *DefaultProductsHttpAdapter) ListProducts(writer http.ResponseWriter, request *http.Request) {

	// TODO implement
}

func (adapter *DefaultProductsHttpAdapter) GetProduct(writer http.ResponseWriter, request *http.Request) {
	// TODO implement
}

func handleErrorResponse(writer http.ResponseWriter, err error) {
	fmt.Println(err)
	writer.WriteHeader(http.StatusBadRequest)
	_, err = writer.Write([]byte(err.Error()))
	if err != nil {
		//TODO log framework
		log.Println(err)
	}
}

func handleResponse(writer http.ResponseWriter, result interface{}, err error) {
	if err != nil {
		handleErrorResponse(writer, err)
		return
	}

	responseBytes, err := json.Marshal(result)
	if err != nil {
		handleErrorResponse(writer, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(responseBytes)
	if err != nil {
		log.Println(err)
	}
}
