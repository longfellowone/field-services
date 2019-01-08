package main

import (
	"log"
	"supply/pkg/mongo"
)

func main() {
	db, err := mongo.NewConnection("default", "password", "supply")
	if err != nil {
		log.Printf("failed to connector to database %v", err)
	}
	service := InitializeOrderingService(db)

	log.Println(service)
}
