package main

import (
	"field/supply/graphql"
	"field/supply/mongo"
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

	productRepository := mongo.NewProductRepository(db)
	searchService, err := search.NewSearchService(productRepository)
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
	}).Handler)

	gqlHandler := graphql.New(searchService)

	router.Handle("/", handler.Playground("Starwars", "/query"))
	router.Handle("/graphql", gqlHandler)

	log.Printf("Listening...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
