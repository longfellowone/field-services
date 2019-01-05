package inmem

import (
	"field/pkg"
	"sync"
)

type OrderRepository struct {
	mu      sync.RWMutex
	orders  map[supply.OrderUUID]*supply.Order
	porders map[supply.ProjectUUID][]*supply.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders:  make(map[supply.OrderUUID]*supply.Order),
		porders: make(map[supply.ProjectUUID][]*supply.Order),
	}
}

func (r *OrderRepository) Save(order *supply.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderUUID] = order

	if _, ok := r.porders[order.ProjectUUID]; !ok {
		r.porders[order.ProjectUUID] = make([]*supply.Order, 0)
	}
	r.porders[order.ProjectUUID] = append(r.porders[order.ProjectUUID], order)

	return nil
}

func (r *OrderRepository) Find(id supply.OrderUUID) (*supply.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if order, ok := r.orders[id]; ok {
		return order, nil
	}
	return nil, supply.ErrOrderNotFound
}

func (r *OrderRepository) FindAllFromProject(id supply.ProjectUUID) ([]*supply.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.porders[id]) == 0 {
		return nil, supply.ErrOrderNotFound
	}
	orders := make([]*supply.Order, 0, len(r.porders))
	for _, o := range r.porders[id] {
		orders = append(orders, o)
	}
	return orders, nil
}
