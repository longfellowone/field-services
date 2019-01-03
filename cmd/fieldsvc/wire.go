//+build wireinject

package main

import (
	"field/pkg"
	"field/pkg/grpc"
	"field/pkg/inmem"
	"github.com/google/wire"
)

//func initializeFieldServices(db *sql.DB) *ordering.Service {
//	wire.Build(
//		postgres.NewOrderRepository,
//		ordering.NewOrderingService,
//		wire.Bind(new(ordering.OrderRepository), &postgres.OrderRepository{}),
//	)
//	return nil
//}

func initializeFieldServicesInMemory() *grpc.Service {
	wire.Build(
		inmem.NewOrderRepository,
		grpc.NewOrderingService,
		wire.Bind(new(material.OrderRepository), &inmem.OrderRepository{}),
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
