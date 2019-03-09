package main

import (
	"context"
	"database/sql"
	"field/supply/graphql"
	"field/supply/ordering"
	"field/supply/postgres"
	"field/supply/search"
	"fmt"
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
	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "default", "password", "localhost", 5432, "default")

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println(err)
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

	srv := &http.Server{Addr: ":8080", Handler: r}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Listening...")

	<-done

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
