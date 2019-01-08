package main

import (
	"context"
	"fmt"
	"log"
	"supply/pkg/ordering"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	var service *ordering.Service

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://default:password@localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}
	//err = client.Disconnect(ctx)

	db := client.Database("supply")
	service = initializeOrderingService(db)

	fmt.Println(service)
}
