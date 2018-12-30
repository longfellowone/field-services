//+build wireinject

package main

import (
	"database/sql"
	"field/pkg/inmem"
	"field/pkg/ordering"
	"field/pkg/postgres"
	"github.com/google/wire"
)

func initializeFieldServices(db *sql.DB) *ordering.Service {
	wire.Build(
		postgres.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(ordering.OrderRepository), &postgres.OrderRepository{}),
	)
	return nil
}

func initializeFieldServicesInMemory() *ordering.Service {
	wire.Build(
		inmem.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(ordering.OrderRepository), &inmem.OrderRepository{}),
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
