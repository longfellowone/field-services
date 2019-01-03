package ordering

import (
	"field/pkg"
	"log"
)

type Service struct {
	db orders.OrderRepository
}

func NewOrderingService(db orders.OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateOrder(o orders.OrderID, p orders.ProjectID) {
	order := orders.NewOrder(o, p)
	err := s.db.Save(order)
	if err != nil {
		log.Println(err)
	}
}

// READ MODELS

func (s *Service) FindOrder(id orders.OrderID) (*orders.Order, error) {
	findAll, err := s.db.Find(id)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}

func (s *Service) FindAllProjectOrders(id orders.ProjectID) ([]*orders.Order, error) {
	findAll, err := s.db.FindAllFromProject(id)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}
