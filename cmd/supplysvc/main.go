package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/mongo"
)

func main() {
	db, err := mongo.Connect("default", "password", "supply")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := InitializeOrderingService(db, s)

	fmt.Println("Listening...")

	// Test
	order := &pb.CreateOrderRequest{
		OrderUuid:   "order1",
		ProjectUuid: "project1",
	}
	_, err = server.CreateOrder(context.TODO(), order)
	if err != nil {
		log.Println(err)
	}
	// End test

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
