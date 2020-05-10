package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/landistas/enterprise-go-examples/cmd/graph"
	"github.com/landistas/enterprise-go-examples/pkg/infra/httpinfra"
)

func _main() error {

	graph, err := graph.Build()
	if err != nil {
		return err
	}

	productsAdapter := graph.ProductsAdapter()
	infraHttpIn := graph.InfraHttpIn()

	err = httpinfra.PublishProducts(infraHttpIn, productsAdapter)
	if err != nil {
		return err
	}
	addr := ":9876"
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      infraHttpIn.MainRouter(),
	}

	fmt.Printf("Listening on %s...", addr)
	return srv.ListenAndServe()
}

func main() {
	if err := _main(); err != nil {
		log.Println(err)
	}
}
