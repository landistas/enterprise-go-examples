package httpinfra

import (
	"net/http"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/inputs/productshttp"
)

func PublishProducts(httpInfraIn *InfraHttpIn, adapter productshttp.ProductsHttpAdapter) error {
	productCreateMatchOptions := HandlerMatchOptions{
		Path:    "/products",
		Methods: []string{http.MethodPost},
	}
	err := httpInfraIn.RegisterRawHandler(productCreateMatchOptions, adapter.CreateProduct)
	if err != nil {
		return err
	}

	productDetailOptions := HandlerMatchOptions{
		Path:    "/products/{productId}",
		Methods: []string{http.MethodGet},
	}
	err = httpInfraIn.RegisterRawHandler(productDetailOptions, adapter.GetProduct)
	if err != nil {
		return err
	}

	productListOptions := HandlerMatchOptions{
		Path:    "/products",
		Methods: []string{http.MethodGet},
	}
	err = httpInfraIn.RegisterRawHandler(productListOptions, adapter.ListProducts)
	if err != nil {
		return err
	}

	return nil
}
