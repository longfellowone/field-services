package main

import (
	"log"
	"supply/pkg/mongo"
)

func main() {
	db, err := mongo.Connect("default", "password", "supply")
	if err != nil {
		log.Printf("failed to connect to database %v", err)
	}
	service := InitializeOrderingService(db)

	log.Println(service)
}
