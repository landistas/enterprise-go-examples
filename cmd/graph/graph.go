package graph

import (
	"github.com/landistas/enterprise-go-examples/pkg/adapters/inputs/productshttp"
	"github.com/landistas/enterprise-go-examples/pkg/adapters/inputs/webhttp"
	"github.com/landistas/enterprise-go-examples/pkg/adapters/outputs/storage"
	"github.com/landistas/enterprise-go-examples/pkg/infra/httpinfra"
	infraStorage "github.com/landistas/enterprise-go-examples/pkg/infra/storage"
	"github.com/landistas/enterprise-go-examples/pkg/usecases"
)

type Graph struct {
	productsAdapter productshttp.ProductsHttpAdapter
	infraHttpIn     *httpinfra.InfraHttpIn
}

func (graph Graph) ProductsAdapter() productshttp.ProductsHttpAdapter {
	if graph.productsAdapter == nil {
		panic("products adapter not initialized")
	}

	return graph.productsAdapter
}

func (graph Graph) InfraHttpIn() *httpinfra.InfraHttpIn {
	if graph.infraHttpIn == nil {
		panic("infraHttpIn not initialized")
	}

	return graph.infraHttpIn
}

func Build() (*Graph, error) {
	graph := &Graph{}
	infraStorage := infraStorage.NewInMemoryInfraStorage()
	productsStorageAdapter := storage.NewProductStorageAdapter(infraStorage)
	productsUseCase := usecases.NewDefaultProductUseCase(productsStorageAdapter)
	productsAdapter := productshttp.NewDefaultProductsHttpAdapter(productsUseCase)

	infraHttpIn := httpinfra.NewInfraHttpIn()
	err := httpinfra.PublishProducts(infraHttpIn, productsAdapter)
	if err != nil {
		return nil, err
	}
	httpinfra.PublishIndex(infraHttpIn, webhttp.NewDefaultWebHttpAdapter())

	graph.productsAdapter = productsAdapter
	graph.infraHttpIn = infraHttpIn

	return graph, nil
}
