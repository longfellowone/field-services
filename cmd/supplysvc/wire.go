//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
	"supply/pkg"
	"supply/pkg/inmem"
	mongodb "supply/pkg/mongo"
	"supply/pkg/ordering"
)

func initializeSupplyService(db *mongo.Database) *ordering.Service {
	wire.Build(
		mongodb.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(supply.OrderRepository), &mongodb.OrderRepository{}),
	)
	return nil
}

func initializeSupplyServiceInMemory() *ordering.Service {
	wire.Build(
		inmem.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(supply.OrderRepository), &inmem.OrderRepository{}),
	)
	return nil
}

//var Set = wire.NewSet(NewOrderRepository)

//func InitializeFieldService() *field.Service {
//	wire.Build(
//		postgres.Set,
//		field.Set,
//	)
//	return nil
//}

//var Set = wire.NewSet(
//	NewOrderRepository,
//	Dial,
//)
