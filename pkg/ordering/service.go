package ordering

import (
	"errors"
	"log"
)

func NewOrderingService(db OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db OrderRepository
}

type OrderRepository interface {
	FindAll() (string, error)
}

func (f *Service) FindAllOrders() (string, error) {
	findAll, err := f.db.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	return findAll, errors.New("not implemented yet")
}
