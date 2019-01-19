package postgres

import (
	"database/sql"
	"errors"
)

//var Set = wire.NewSet(NewOrderRepository)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (c *OrderRepository) FindAll() (string, error) {
	return "", errors.New("FindAll() not implemented")
}
