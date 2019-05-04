package main

import (
	"database/sql"
	server "field/supply/grpc"
	"field/supply/ordering"
	"field/supply/postgres"
	"field/supply/search"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	dbConfig := postgres.Config{
		DBHost:     "35.197.72.107",
		DBPort:     5432,
		DBUser:     "postgres",
		DBPassword: "uR57xGK6k7Hq",
		DBName:     "postgres",
	}

	db, err := postgres.Connect(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Listening on :9090")

	s := InitializeSupplyServices(db)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitializeSupplyServices(db *sql.DB) *grpc.Server {
	orderRepository := postgres.NewOrderRepository(db)
	productRepository := postgres.NewProductRepository(db)
	projectRepository := postgres.NewProjectRepository(db)

	orderingService := ordering.NewOrderingService(orderRepository, projectRepository)
	searchService := search.NewSearchService(productRepository)

	return server.New(orderingService, searchService)
}
