package main

import (
	"field/supply/graphql"
	"field/supply/mongo"
	"field/supply/search"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := mongo.Connect("default", "password", "supply")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	productRepository := mongo.NewProductRepository(db)
	searchService, err := search.NewSearchService(productRepository)
	if err != nil {
		log.Fatal(err)
	}

	server := graphql.New(searchService)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
