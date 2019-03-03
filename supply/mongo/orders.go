package mongo

import (
	"context"
	"field/supply"
	"field/supply/ordering"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

type OrderRepository struct {
	coll *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	coll := db.Collection("orders")

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	orderIndex := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "orderid", Value: bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	projectIndex := mongo.IndexModel{
		Keys: bsonx.Doc{{Key: "projectid", Value: bsonx.Int32(1)}},
	}
	indexModel := []mongo.IndexModel{orderIndex, projectIndex}

	_, err := coll.Indexes().CreateMany(context.TODO(), indexModel, opts)
	if err != nil {
		log.Fatal(err)
	}

	return &OrderRepository{
		coll: coll,
	}
}

func (r *OrderRepository) Save(o *supply.Order) error {
	order, err := bson.Marshal(&o)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "orderid", Value: o.OrderID}}
	opts := options.Replace().SetUpsert(true)

	_, err = r.coll.ReplaceOne(context.TODO(), filter, order, opts)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *OrderRepository) Find(id string) (*supply.Order, error) {
	var order supply.Order

	filter := bson.D{{Key: "orderid", Value: id}}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindDates(projectid string) ([]ordering.ProjectOrder, error) {
	var orders []ordering.ProjectOrder
	filter := bson.D{{Key: "projectid", Value: projectid}}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var order ordering.ProjectOrder
		err := cur.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
