// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"field/pkg/inmem"
	"field/pkg/ordering"
)

// Injectors from wire.go:

func initializeFieldServicesInMemory() *ordering.Service {
	orderRepository := inmem.NewOrderRepository()
	service := ordering.NewOrderingService(orderRepository)
	return service
}