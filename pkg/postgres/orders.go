package postgres

import "database/sql"

//var Set = wire.NewSet(NewOrderRepository)

type OrderRepository struct{}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (c *OrderRepository) FindAll() (string, error) {
	return "Found All!", nil
}
