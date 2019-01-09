package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
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
	InitializeOrderingService(db, s)

	fmt.Println("Listening...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
