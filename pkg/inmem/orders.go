package inmem

import (
	"errors"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (c *OrderRepository) FindAll() (string, error) {
	return "", errors.New("FindAll() not implemented")
}
