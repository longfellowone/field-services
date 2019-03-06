package mongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

// Add disconnect function and call in main
// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/getting-started/bookshelf/db_mongo.go

//db, err := mongo.Connect("default", "password", "supply")
//if err != nil {
//	log.Fatalf("failed to connect to database: %v", err)
//}
//orderRepository := mongo.NewOrderRepository(db)
//productRepository := mongo.NewProductRepository(db)

func Connect(username, passwrd, dbname string) (*mongo.Database, error) {
	conn := fmt.Sprintf("mongodb://%s:%s@localhost:27017", username, passwrd)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, conn)
	if err != nil {
		return &mongo.Database{}, err
	}
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return &mongo.Database{}, err
	}
	//err = client.Disconnect(ctx)

	return client.Database(dbname), nil
}
