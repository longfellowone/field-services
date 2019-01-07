package mongo

import (
	"context"
	"log"
	supply "supply/pkg"
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

	// Update after next release, see options.IndexOptions struct
	// https://godoc.org/github.com/mongodb/mongo-go-driver/mongo/options#IndexOptions
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	orderIndex := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "orderuuid", Value: bsonx.Int32(1)}},
		Options: bsonx.Doc{{Key: "unique", Value: bsonx.Boolean(true)}},
	}
	projectIndex := mongo.IndexModel{
		Keys: bsonx.Doc{{Key: "projectuuid", Value: bsonx.Int32(1)}},
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
	//fmt.Println(o)

	order, err := bson.Marshal(&o)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "orderuuid", Value: o.OrderUUID}}
	opts := options.Replace().SetUpsert(true)

	_, err = r.db.ReplaceOne(context.TODO(), filter, order, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Find(uuid supply.OrderUUID) (*supply.Order, error) {
	var order supply.Order

	filter := bson.D{{Key: "orderuuid", Value: uuid}}

	err := r.db.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindAllFromProject(uuid supply.ProjectUUID) ([]supply.Order, error) {
	var orders []supply.Order
	filter := bson.D{{Key: "projectuuid", Value: uuid}}

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
