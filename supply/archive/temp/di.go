// https://play.golang.org/p/Nn3Wv0TwkMe
package main

import "fmt"

func main() {
	db := NewOrderRespository() // Create repository
	service := NewService(db)   // Inject into service

	service.CreateNewOrder("order1") // Create sample order
	service.FindOrder("order1")      // Find the order just created
}

// Service
type orderRepository interface {
	Find(id string) *Order
	Save(o *Order)
}

type Service struct {
	db orderRepository
}

func NewService(db orderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateNewOrder(id string) {
	order := CreateNewOrder(id)
	s.db.Save(order)
	fmt.Println("Saved:", order.OrderID)
}

func (s *Service) FindOrder(id string) {
	order := s.db.Find(id)
	fmt.Println("Found:", order.OrderID)
}

// Repository
type OrderRepository struct {
	order map[string]*Order
}

func NewOrderRespository() *OrderRepository {
	order := make(map[string]*Order)
	return &OrderRepository{
		order: order,
	}
}

func (r *OrderRepository) Find(id string) *Order {
	order := r.order[id]
	return order
}

func (r *OrderRepository) Save(o *Order) {
	r.order[o.OrderID] = o
}

// Domain
type Order struct {
	OrderID string
}

func CreateNewOrder(id string) *Order {
	return &Order{
		OrderID: id,
	}
}
