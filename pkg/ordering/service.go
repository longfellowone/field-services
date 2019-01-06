package ordering

import (
	"field/pkg"
	"log"
)

type Service struct {
	db supply.OrderRepository
}

func NewOrderingService(db supply.OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateOrder(o supply.OrderUUID, p supply.ProjectUUID) {
	order := supply.Create(o, p)
	err := s.db.Save(order)
	if err != nil {
		log.Println(err)
	}
}

func (s *Service) FindOrder(uuid supply.OrderUUID) (*supply.Order, error) {
	findAll, err := s.db.Find(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}

func (s *Service) FindAllProjectOrders(uuid supply.ProjectUUID) ([]*supply.Order, error) {
	findAll, err := s.db.FindAllFromProject(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}
