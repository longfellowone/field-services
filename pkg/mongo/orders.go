package mongo

import (
	"field/pkg"
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

func (r *OrderRepository) Save(order *supply.Order) error {

	return errors.New("Not Implemented")
}

func (r *OrderRepository) Find(uuid supply.OrderUUID) (*supply.Order, error) {
	return nil, errors.New("Not Implemented")
}

func (r *OrderRepository) FindAllFromProject(uuid supply.ProjectUUID) ([]*supply.Order, error) {
	return nil, errors.New("Not Implemented")
}
