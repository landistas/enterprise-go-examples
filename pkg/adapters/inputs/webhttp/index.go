package webhttp

import (
	"fmt"
	"html/template"
	"net/http"
)

type WebHttpAdapter interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type DefaultProductsHttpAdapter struct{}

func NewDefaultWebHttpAdapter() *DefaultProductsHttpAdapter {
	return &DefaultProductsHttpAdapter{}
}

func (adapter *DefaultProductsHttpAdapter) Index(response http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, err) // TODO: log it
		return
	}
	err = t.Execute(response, nil)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, err)
		return
	}
}
