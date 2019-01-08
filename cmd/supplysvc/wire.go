//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"supply/pkg"
	mongodb "supply/pkg/mongo"
	"supply/pkg/ordering"
)

func initializeOrderingService(db *mongo.Database) *ordering.Service {
	wire.Build(
		mongodb.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(supply.OrderRepository), &mongodb.OrderRepository{}),
	)
	return nil
}
