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
	ReceiveOrderItem(orderid, productid string, quantity uint) error
}

type Service struct {
	order supply.OrderRepository
}

func NewOrderingService(order supply.OrderRepository) *Service {
	return &Service{
		order: order,
	}
}

func (s *Service) CreateOrder(orderid, projectid string) error {
	order := supply.Create(orderid, projectid)

	err := s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddOrderItem(orderid, productid, name, uom string) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	err = order.AddItem(productid, name, uom)
	if err != nil {
		return err
	}

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveOrderItem(orderid, productid string) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	err = order.RemoveItem(productid)
	if err != nil {
		return err
	}

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ModifyRequestedQuantity(orderid, productid string, quantity uint) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	err = order.UpdateQuantityRequested(productid, quantity)
	if err != nil {
		return err
	}

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendOrder(orderid string) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	err = order.Send()
	if err != nil {
		return err
	}

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReceiveOrderItem(orderid, productid string, quantity uint) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	err = order.ReceiveItem(productid, quantity)
	if err != nil {
		return err
	}

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FindAllProjectOrders(uuid string) ([]supply.Order, error) {
	findAll, err := s.order.FindAllFromProject(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}

func (s *Service) FindOrder(uuid string) (*supply.Order, error) {
	findAll, err := s.order.Find(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}
