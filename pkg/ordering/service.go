package ordering

import (
	"log"
	"supply/pkg"
)

type OrderingService interface {
	CreateOrder(orderid, projectid string) error
	AddOrderItem(orderid, productid, name, uom string) error
	RemoveOrderItem(orderid, productid string) error
	ModifyRequestedQuantity(orderid, productid string, quantity uint) error
	SendOrder(orderid string) error
	UpdateItemPO(orderid, productid, ponumber string) error
	ReceiveOrderItem(orderid, productid string, quantity uint) error
}

type Service struct {
	db supply.OrderRepository
}

func NewOrderingService(db supply.OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateOrder(orderid, projectid string) error {
	order := supply.Create(orderid, projectid)

	err := s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddOrderItem(orderid, productid, name, uom string) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	err = order.AddItem(productid, name, uom)
	if err != nil {
		return err
	}

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveOrderItem(orderid, productid string) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	err = order.RemoveItem(productid)
	if err != nil {
		return err
	}

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ModifyRequestedQuantity(orderid, productid string, quantity uint) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	err = order.UpdateQuantityRequested(productid, quantity)
	if err != nil {
		return err
	}

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendOrder(orderid string) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	err = order.Send()
	if err != nil {
		return err
	}

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateItemPO(orderid, productid, ponumber string) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	err = order.UpdatePO(productid, ponumber)
	if err != nil {
		return err
	}

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReceiveOrderItem(orderid, productid string, quantity uint) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	err = order.ReceiveItem(productid, quantity)
	if err != nil {
		return err
	}

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FindAllProjectOrders(uuid string) ([]supply.Order, error) {
	findAll, err := s.db.FindAllFromProject(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}

func (s *Service) FindOrder(uuid string) (*supply.Order, error) {
	findAll, err := s.db.Find(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}
