//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"supply/pkg/grpc"
	mongodb "supply/pkg/mongo"
	"supply/pkg/ordering"
)

func InitializeOrderingServices(db *mongo.Database) *server.OrderingServer {
	wire.Build(
		mongodb.NewOrderRepository,
		wire.Bind(new(ordering.OrderRepository), &mongodb.OrderRepository{}),
		ordering.NewOrderingService,
		wire.Bind(new(ordering.OrderingService), &ordering.Service{}),
		server.NewOrderingServer,
	)
	return nil
}

//func InitializePurchasingServices(db *mongo.Database, svr *grpc.Server) *server.PurchasingServer {
//	wire.Build(
//		mongodb.NewOrderRepository,
//		mongodb.NewProductRepository,
//		wire.Bind(new(purchasing.OrderRepository), &mongodb.OrderRepository{}),
//		wire.Bind(new(purchasing.ProductRepository), &mongodb.ProductRepository{}),
//		purchasing.NewPurchasingService,
//		wire.Bind(new(purchasing.PurchasingService), &purchasing.Service{}),
//		server.NewPurchasingServer,
//	)
//	return nil
//}
