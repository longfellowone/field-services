package postgres

import (
	"database/sql"
	"field/supply"
	"field/supply/ordering"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Save(o *supply.Order) error {
	return nil
}

func (r *OrderRepository) Find(id string) (*supply.Order, error) {
	var order supply.Order

	return &order, nil
}

func (r *OrderRepository) FindDates(projectid string) ([]ordering.ProjectOrder, error) {
	var orders []ordering.ProjectOrder

	return orders, nil
}
