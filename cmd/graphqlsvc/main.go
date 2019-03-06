package main

import (
	"field/supply/graphql"
	"field/supply/ordering"
	"field/supply/postgres"
	"field/supply/search"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.Connect("localhost", 5432, "default", "password", "default")
	if err != nil {
		panic(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	productRepository := postgres.NewProductRepository(db)

	orderingService := ordering.NewOrderingService(orderRepository)
	searchService := search.NewSearchService(productRepository)

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler)

	gqlHandler := graphql.Initialize(searchService, orderingService)

	r.Handle("/graphql", gqlHandler)
	r.Handle("/", handler.Playground("", "/graphql"))

	log.Printf("Listening...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
