package httpinfra

import (
	"net/http"

	"github.com/gorilla/mux"
)

type InfraHttpIn struct {
	mainRouter *mux.Router
}

func (infraHttpIn *InfraHttpIn) MainRouter() *mux.Router {
	return infraHttpIn.mainRouter
}

func NewInfraHttpIn() *InfraHttpIn {
	router := mux.NewRouter()
	return &InfraHttpIn{
		mainRouter: router,
	}
}

type HandlerMatchOptions struct {
	Path    string
	Methods []string
}

type RawHandler func(w http.ResponseWriter, r *http.Request)

type RawMiddleware func(http.Handler) http.Handler

func (infraHttpIn *InfraHttpIn) RegisterRawHandler(matchOptions HandlerMatchOptions, handler RawHandler) error {
	infraHttpIn.mainRouter.HandleFunc(matchOptions.Path, handler).Methods(matchOptions.Methods...)
	return nil
}

func (infraHttpIn *InfraHttpIn) RegisterMiddleware(middleware RawMiddleware) error {
	infraHttpIn.mainRouter.Use(mux.MiddlewareFunc(middleware))
	return nil
}
