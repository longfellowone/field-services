package inmem

import (
	"field/pkg"
	"sync"
)

type OrderRepository struct {
	mu      sync.RWMutex
	orders  map[field.OrderUUID]*field.Order
	porders map[field.ProjectUUID][]*field.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders:  make(map[field.OrderUUID]*field.Order),
		porders: make(map[field.ProjectUUID][]*field.Order),
	}
}

func (r *OrderRepository) Save(order *field.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderUUID] = order

	if _, ok := r.porders[order.ProjectUUID]; !ok {
		r.porders[order.ProjectUUID] = make([]*field.Order, 0)
	}
	r.porders[order.ProjectUUID] = append(r.porders[order.ProjectUUID], order)

	return nil
}

func (r *OrderRepository) Find(id field.OrderUUID) (*field.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if order, ok := r.orders[id]; ok {
		return order, nil
	}
	return nil, field.ErrOrderNotFound
}

func (r *OrderRepository) FindAllFromProject(id field.ProjectUUID) ([]*field.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.porders[id]) == 0 {
		return nil, field.ErrOrderNotFound
	}
	orders := make([]*field.Order, 0, len(r.porders))
	for _, o := range r.porders[id] {
		orders = append(orders, o)
	}
	return orders, nil
}
