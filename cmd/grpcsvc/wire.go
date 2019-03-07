//+build wireinject
// BROKEN -> TO FIX, EXPORT REPOSITORY INTERFACES AND SERVICE TYPES

package main

import (
	"field/supply/grpc"
	mongodb "field/supply/mongo"
	"field/supply/ordering"
	"field/supply/purchasing"
	"field/supply/search"
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"google.golang.org/grpc"
)

func InitializeSupplyServices(db *mongo.Database) (*grpc.Server, error) {
	wire.Build(
		// Repositories
		mongodb.NewOrderRepository,
		mongodb.NewProductRepository,
		// Ordering service
		wire.Bind(new(ordering.orderRepository), &mongodb.OrderRepository{}),
		ordering.NewOrderingService,
		wire.Bind(new(ordering.Service), &ordering.service{}),
		// Purchasing service
		wire.Bind(new(purchasing.productRepository), &mongodb.ProductRepository{}),
		wire.Bind(new(purchasing.orderRepository), &mongodb.OrderRepository{}),
		purchasing.NewPurchasingService,
		wire.Bind(new(purchasing.Service), &purchasing.service{}),
		// Search service
		wire.Bind(new(search.productRepository), &mongodb.ProductRepository{}),
		search.NewSearchService,
		wire.Bind(new(search.Service), &search.service{}),
		// gRPC server
		server.New,
	)
	return &grpc.Server{}, nil
}
