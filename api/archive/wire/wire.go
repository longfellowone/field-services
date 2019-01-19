package wire

import (
	"github.com/google/wire"
	"supply/api"
	"supply/api/ordering"
)

func initializeSupplyServiceInMemory() *ordering.Service {
	wire.Build(
		inmem.NewOrderRepository,
		ordering.NewOrderingService,
		wire.Bind(new(supply.OrderRepository), &inmem.OrderRepository{}),
	)
	return nil
}

var Set = wire.NewSet(NewOrderRepository)

func InitializeFieldService() *field.Service {
	wire.Build(
		postgres.Set,
		field.Set,
	)
	return nil
}

var Set = wire.NewSet(
	NewOrderRepository,
	Dial,
)
