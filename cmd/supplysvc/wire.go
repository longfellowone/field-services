//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"google.golang.org/grpc"
	"supply/pkg/grpc"
	mongodb "supply/pkg/mongo"
	"supply/pkg/ordering"
)

func InitializeSupplyServices(db *mongo.Database, svr *grpc.Server) *server.OrderingServer {
	wire.Build(
		mongodb.NewOrderRepository,
		wire.Bind(new(ordering.OrderRepository), &mongodb.OrderRepository{}),
		ordering.NewOrderingService,
		wire.Bind(new(ordering.OrderingService), &ordering.Service{}),
		server.NewOrderingServer,
	)
	return nil
}
