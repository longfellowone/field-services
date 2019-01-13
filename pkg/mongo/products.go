package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"log"
	"supply/pkg"
	"time"
)

type ProductRepository struct {
	coll  *mongo.Collection
	cache []supply.Product
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	coll := db.Collection("products")

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	indexModel := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "productid", Value: bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}

	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel, opts)
	if err != nil {
		log.Fatal(err)
	}

	return &ProductRepository{
		coll: coll,
	}
}

func (r *ProductRepository) Save(p *supply.Product) error {
	product, err := bson.Marshal(&p)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "productid", Value: p.ProductID}}
	opts := options.Replace().SetUpsert(true)

	_, err = r.coll.ReplaceOne(context.TODO(), filter, product, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Find(id string) (*supply.Product, error) {
	var product supply.Product

	filter := bson.D{{Key: "productid", Value: id}}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindAll() ([]supply.Product, error) {
	var products []supply.Product
	filter := bson.D{{}}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var product supply.Product
		err := cur.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
