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
)

func InitializeOrderingServices(db *mongo.Database) (*grpc.Server, error) {
	wire.Build(
		// Ordering service
		mongodb.NewOrderRepository,
		wire.Bind(new(ordering.OrderRepository), &mongodb.OrderRepository{}),
		ordering.NewOrderingService,
		wire.Bind(new(ordering.OrderingService), &ordering.Service{}),
		// Purchasing service
		mongodb.NewProductRepository,
		wire.Bind(new(purchasing.ProductRepository), &mongodb.ProductRepository{}),
		wire.Bind(new(purchasing.OrderRepository), &mongodb.OrderRepository{}),
		purchasing.NewPurchasingService,
		wire.Bind(new(purchasing.PurchasingService), &purchasing.Service{}),
		// gRPC server
		server.New,
	)
	return &grpc.Server{}, nil
}
