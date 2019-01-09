//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"google.golang.org/grpc"
	"supply/pkg"
	server "supply/pkg/grpc"
	mongodb "supply/pkg/mongo"
	"supply/pkg/ordering"
)

func InitializeOrderingService(db *mongo.Database, svr *grpc.Server) *server.Server {
	wire.Build(
		mongodb.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(supply.OrderRepository), &mongodb.OrderRepository{}),
		server.NewOrderingServer,
		wire.Bind(new(server.OrderingService), &ordering.Service{}),
	)
	return nil
}
