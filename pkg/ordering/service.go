package ordering

import (
	"supply/pkg"
	"time"
)

type OrderingService interface {
	CreateOrder(orderid, projectid string) error
	AddOrderItem(orderid, productid, name, uom string) error
	RemoveOrderItem(orderid, productid string) error
	ModifyRequestedQuantity(orderid, productid string, quantity uint) error
	SendOrder(orderid string) error
	ReceiveOrderItem(orderid, productid string, quantity uint) error
	FindOrder(orderid string) (supply.Order, error)
	QueryOrdersFromProject(projectid string) ([]ProjectOrder, error)
}

type OrderRepository interface {
	supply.OrderRepository
	QueryOrdersFromProject(projectid string) ([]ProjectOrder, error)
}

type Service struct {
	order OrderRepository
}

func NewOrderingService(order OrderRepository) *Service {
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

func (s *Service) FindOrder(orderid string) (supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return supply.Order{}, err
	}
	return *order, nil
}

// Order read models

type ProjectOrder struct {
	OrderDate time.Time
}

func (s *Service) QueryOrdersFromProject(projectid string) ([]ProjectOrder, error) {
	order, err := s.order.QueryOrdersFromProject(projectid)
	if err != nil {
		return []ProjectOrder{}, err
	}
	return order, nil
}
