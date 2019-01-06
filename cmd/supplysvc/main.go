package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"supply/pkg/ordering"
	"time"
)

const (
	inMemory = false
)

func main() {
	var service *ordering.Service

	if inMemory {
		service = initializeSupplyServiceInMemory()
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		client, err := mongo.Connect(ctx, "mongodb://default:password@localhost:27017")
		if err != nil {
			log.Fatal(err)
		}
		defer cancel()

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		//err = client.Disconnect(ctx)

		db := client.Database("field")
		service = initializeSupplyService(db)
	}

	fmt.Println(service)
}
