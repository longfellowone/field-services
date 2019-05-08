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
	"os"
)

const (
	defaultDBHost     = "localhost"
	defaultDBUser     = "default"
	defaultDBPassword = "password"
	defaultDBName     = "default"
)

func main() {
	dbConfig := postgres.Config{
		DBHost:     envString("DB_HOSTNAME", defaultDBHost),
		DBPort:     5432,
		DBUser:     envString("DB_USER", defaultDBUser),
		DBPassword: envString("DB_PASSWORD", defaultDBPassword),
		DBName:     envString("DB_NAME", defaultDBName),
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

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
