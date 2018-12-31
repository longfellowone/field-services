package inmem

import (
	"field/pkg"
	"sync"
)

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[material.OrderID]*material.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[material.OrderID]*material.Order),
	}
}

func (r *OrderRepository) Save(o *material.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[o.OrderID] = o
	return nil
}

func (r *OrderRepository) Find(id material.OrderID) (*material.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if o, ok := r.orders[id]; ok {
		return o, nil
	}
	return nil, material.ErrOrderNotFound
}

//func (r *OrderRepository) FindAll() (string, error) {
//	return "", errors.New("not implemented")
//}

//func (r *OrderRepository) FindAll() (string, error) {
//	return "", errors.New("not implemented")
//}
