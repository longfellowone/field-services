package mongo

import (
	"context"
	"field/pkg"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

type OrderRepository struct {
	db *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	collection := db.Collection("orders")

	return &OrderRepository{
		db: collection,
	}
}

func (r *OrderRepository) Save(o *supply.Order) error {
	order, err := bson.Marshal(&o)
	if err != nil {
		return err
	}

	_, err = r.db.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Find(uuid supply.OrderUUID) (*supply.Order, error) {
	var order supply.Order

	err := r.db.FindOne(context.TODO(), nil).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindAllFromProject(uuid supply.ProjectUUID) ([]*supply.Order, error) {

	return nil, errors.New("Not Implemented")
}
