package ordering

import (
	"field/pkg"
	"log"
)

type Service struct {
	db material.OrderRepository
}

func NewOrderingService(db material.OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateNewOrder(o material.OrderID, p material.ProjectID) (*material.Order, error) {
	order, err := material.NewOrder(o, p)
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.Save(order)
	if err != nil {
		log.Fatal(err)
	}
	return order, nil
}

func (s *Service) FindOrder(id material.OrderID) (*material.Order, error) {
	findAll, err := s.db.Find(id)
	if err != nil {
		log.Fatal(err)
	}
	return findAll, nil
}

func (s *Service) FindAllProjectOrders(id material.ProjectID) ([]*material.Order, error) {
	findAll, err := s.db.FindAllFromProject(id)
	if err != nil {
		log.Fatal(err)
	}
	return findAll, nil
}
