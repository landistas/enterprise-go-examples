package httpinfra

import (
	"net/http"

	"github.com/landistas/enterprise-go-examples/pkg/adapters/inputs/webhttp"
)

func PublishIndex(httpInfraIn *InfraHttpIn, adapter webhttp.WebHttpAdapter) {

	IndexMatchOptions := HandlerMatchOptions{
		Path:    "/",
		Methods: []string{http.MethodGet},
	}
	err := httpInfraIn.RegisterRawHandler(IndexMatchOptions, adapter.Index)
	if err != nil {
		panic(err)
	}
}
