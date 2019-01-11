// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"supply/pkg/grpc"
	mongo2 "supply/pkg/mongo"
	"supply/pkg/ordering"
)

// Injectors from wire.go:

func InitializeOrderingServices(db *mongo.Database) *server.OrderingServer {
	orderRepository := mongo2.NewOrderRepository(db)
	service := ordering.NewOrderingService(orderRepository)
	orderingServer := server.NewOrderingServer(service)
	return orderingServer
}
