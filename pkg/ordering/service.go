package ordering

import (
	"field/pkg"
	"log"
)

type Service struct {
	db field.OrderRepository
}

func NewOrderingService(db field.OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateOrder(o field.OrderUUID, p field.ProjectUUID) {
	order := field.Create(o, p)
	err := s.db.Save(order)
	if err != nil {
		log.Println(err)
	}
}

// READ MODELS

func (s *Service) FindOrder(id field.OrderUUID) (*field.Order, error) {
	findAll, err := s.db.Find(id)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}

func (s *Service) FindAllProjectOrders(id field.ProjectUUID) ([]*field.Order, error) {
	findAll, err := s.db.FindAllFromProject(id)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}
