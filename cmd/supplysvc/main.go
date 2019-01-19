package main

import (
	"fmt"
	"log"
	"net"
	"supply/api/mongo"
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

	s, err := InitializeSupplyServices(db)
	if err != nil {
		log.Fatalf("failed to initialize supply services: %v", err)
	}

	fmt.Println("Listening...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
