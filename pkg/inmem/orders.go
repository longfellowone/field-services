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

func (r *OrderRepository) Save(o *material.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[o.OrderID] = o

	if _, ok := r.porders[o.ProjectID]; !ok {
		r.porders[o.ProjectID] = make([]*material.Order, 0)
	}
	r.porders[o.ProjectID] = append(r.porders[o.ProjectID], o)

	return nil
}

func (r *OrderRepository) Find(id material.OrderID) (*material.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if val, ok := r.orders[id]; ok {
		return val, nil
	}
	return nil, material.ErrOrderNotFound
}

func (r *OrderRepository) FindAllFromProject(id material.ProjectID) ([]*material.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.porders[id]) == 0 {
		return nil, material.ErrOrderNotFound
	}
	o := make([]*material.Order, 0, len(r.porders))
	for _, v := range r.porders[id] {
		o = append(o, v)
	}
	return o, nil
}

//func (r *OrderRepository) FindAll() (string, error) {
//	return "", errors.New("not implemented")
//}

//func (r *OrderRepository) FindAll() (string, error) {
//	return "", errors.New("not implemented")
//}
