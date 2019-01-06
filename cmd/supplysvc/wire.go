//+build wireinject

package main

import (
	"field/pkg"
	"field/pkg/inmem"
	mongodb "field/pkg/mongo"
	"field/pkg/ordering"
	"github.com/google/wire"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func initializeFieldServices(db *mongo.Database) *ordering.Service {
	wire.Build(
		mongodb.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(supply.OrderRepository), &mongodb.OrderRepository{}),
	)
	return nil
}

func initializeFieldServicesInMemory() *ordering.Service {
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
