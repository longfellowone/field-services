package main

import (
	"context"
	"field/supply/graphql"
	"field/supply/ordering"
	"field/supply/postgres"
	"field/supply/search"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := postgres.Connect("localhost", 5432, "default", "password", "default")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	srv := &http.Server{Addr: ":8080", Handler: r}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Listening...")

	<-done

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
