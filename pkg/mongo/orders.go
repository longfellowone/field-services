package mongo

import (
	"context"
	"log"
	"supply/pkg"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

type OrderRepository struct {
	db *mongo.Collection
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
		db: coll,
	}
}

func (r *OrderRepository) Save(o *supply.Order) error {
	order, err := bson.Marshal(&o)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "orderid", Value: o.OrderID}}
	opts := options.Replace().SetUpsert(true)

	_, err = r.db.ReplaceOne(context.TODO(), filter, order, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Find(id string) (*supply.Order, error) {
	var order supply.Order

	filter := bson.D{{Key: "orderid", Value: id}}

	err := r.db.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindAllFromProject(id string) ([]supply.Order, error) {
	var orders []supply.Order
	filter := bson.D{{Key: "projectid", Value: id}}

	cur, err := r.db.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var order supply.Order
		err := cur.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return orders, nil
}
