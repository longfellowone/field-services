package inmem

import (
	"field/pkg"
	"sync"
)

type OrderRepository struct {
	mu      sync.RWMutex
	orders  map[material.OrderID]*material.Order
	porders map[material.ProjectID][]*material.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders:  make(map[material.OrderID]*material.Order),
		porders: make(map[material.ProjectID][]*material.Order),
	}
}

func (r *OrderRepository) Save(order *material.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderID] = order

	if _, ok := r.porders[order.ProjectID]; !ok {
		r.porders[order.ProjectID] = make([]*material.Order, 0)
	}
	r.porders[order.ProjectID] = append(r.porders[order.ProjectID], order)

	return nil
}

func (r *OrderRepository) Find(id material.OrderID) (*material.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if order, ok := r.orders[id]; ok {
		return order, nil
	}
	return nil, material.ErrOrderNotFound
}

func (r *OrderRepository) FindAllFromProject(id material.ProjectID) ([]*material.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.porders[id]) == 0 {
		return nil, material.ErrOrderNotFound
	}
	orders := make([]*material.Order, 0, len(r.porders))
	for _, o := range r.porders[id] {
		orders = append(orders, o)
	}
	return orders, nil
}
