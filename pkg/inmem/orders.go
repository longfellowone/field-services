package inmem

import (
	"field/pkg"
	"sync"
)

type OrderRepository struct {
	mu      sync.RWMutex
	orders  map[orders.OrderID]*orders.Order
	porders map[orders.ProjectID][]*orders.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders:  make(map[orders.OrderID]*orders.Order),
		porders: make(map[orders.ProjectID][]*orders.Order),
	}
}

func (r *OrderRepository) Save(order *orders.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderID] = order

	if _, ok := r.porders[order.ProjectID]; !ok {
		r.porders[order.ProjectID] = make([]*orders.Order, 0)
	}
	r.porders[order.ProjectID] = append(r.porders[order.ProjectID], order)

	return nil
}

func (r *OrderRepository) Find(id orders.OrderID) (*orders.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if order, ok := r.orders[id]; ok {
		return order, nil
	}
	return nil, orders.ErrOrderNotFound
}

func (r *OrderRepository) FindAllFromProject(id orders.ProjectID) ([]*orders.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.porders[id]) == 0 {
		return nil, orders.ErrOrderNotFound
	}
	orders := make([]*orders.Order, 0, len(r.porders))
	for _, o := range r.porders[id] {
		orders = append(orders, o)
	}
	return orders, nil
}
