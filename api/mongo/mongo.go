package mongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

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
