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
	dbConfig := postgres.Config{
		DBHost:     "localhost",
		DBPort:     5432,
		DBUser:     "default",
		DBPassword: "password",
		DBName:     "default",
	}

	db, err := postgres.Connect(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderRepository := postgres.NewOrderRepository(db)
	productRepository := postgres.NewProductRepository(db)
	projectRepository := postgres.NewProjectRepository(db)

	orderingService := ordering.NewOrderingService(orderRepository, projectRepository)
	searchService := search.NewSearchService(productRepository)

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler)

	// Uncomment to use Auth0 middleware
	//r.Use(auth.Middleware())

	gqlHandler := graphql.Initialize(searchService, orderingService)

	r.Handle("/graphqlsvc", gqlHandler)
	r.Handle("/", handler.Playground("", "/graphqlsvc"))

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
