//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"google.golang.org/grpc"
	"supply/pkg/grpc"
	mongodb "supply/pkg/mongo"
	"supply/pkg/ordering"
	"supply/pkg/purchasing"
	"supply/pkg/search"
)

func InitializeOrderingServices(db *mongo.Database) (*grpc.Server, error) {
	wire.Build(
		// Repositories
		mongodb.NewOrderRepository,
		mongodb.NewProductRepository,
		// Ordering service
		wire.Bind(new(ordering.OrderRepository), &mongodb.OrderRepository{}),
		ordering.NewOrderingService,
		wire.Bind(new(ordering.OrderingService), &ordering.Service{}),
		// Purchasing service
		wire.Bind(new(purchasing.ProductRepository), &mongodb.ProductRepository{}),
		wire.Bind(new(purchasing.OrderRepository), &mongodb.OrderRepository{}),
		purchasing.NewPurchasingService,
		wire.Bind(new(purchasing.PurchasingService), &purchasing.Service{}),
		// Search service
		wire.Bind(new(search.ProductRepository), &mongodb.ProductRepository{}),
		search.NewSearchService,
		wire.Bind(new(search.SearchService), &search.Service{}),
		// gRPC server
		server.New,
	)
	return &grpc.Server{}, nil
}
