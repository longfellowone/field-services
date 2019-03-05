package main

import (
	"field/supply/graphql"
	"field/supply/mongo"
	"field/supply/ordering"
	"field/supply/search"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	db, err := mongo.Connect("default", "password", "supply")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	orderRepository := mongo.NewOrderRepository(db)
	productRepository := mongo.NewProductRepository(db)

	orderingService := ordering.NewOrderingService(orderRepository)
	searchService := search.NewSearchService(productRepository)

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
	}).Handler)

	gqlHandler := graphql.New(searchService, orderingService)

	r.Handle("/", handler.Playground("", "/graphql"))
	r.Handle("/graphql", gqlHandler)

	log.Printf("Listening...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
