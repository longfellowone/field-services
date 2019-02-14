package inmem

//
//import (
//	"supply/supply"
//	"sync"
//)
//
//type OrderRepository struct {
//	mu      sync.RWMutex
//	orders  map[string]*supply.Order
//	porders map[string][]*supply.Order
//}
//
//func NewOrderRepository() *OrderRepository {
//	return &OrderRepository{
//		orders:  make(map[string]*supply.Order),
//		porders: make(map[string][]*supply.Order),
//	}
//}
//
//func (r *OrderRepository) Save(order *supply.Order) error {
//	r.mu.Lock()
//	defer r.mu.Unlock()
//
//	r.orders[order.OrderUUID] = order
//
//	if _, ok := r.porders[order.ProjectUUID]; !ok {
//		r.porders[order.ProjectUUID] = make([]*supply.Order, 0)
//	}
//	r.porders[order.ProjectUUID] = append(r.porders[order.ProjectUUID], order)
//
//	return nil
//}
//
//func (r *OrderRepository) Find(uuid string) (*supply.Order, error) {
//	r.mu.RLock()
//	defer r.mu.RUnlock()
//
//	if order, ok := r.orders[uuid]; ok {
//		return order, nil
//	}
//	return nil, supply.ErrOrderNotFound
//}
//
//func (r *OrderRepository) FindAllFromProject(uuid string) ([]supply.Order, error) {
//	r.mu.RLock()
//	defer r.mu.RUnlock()
//
//	if len(r.porders[uuid]) == 0 {
//		return nil, supply.ErrOrderNotFound
//	}
//	orders := make([]supply.Order, 0, len(r.porders))
//	for _, o := range r.porders[uuid] {
//		orders = append(orders, *o)
//	}
//	return orders, nil
//}
