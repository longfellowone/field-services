package ordering

import (
	"field/supply"
)

type Service interface {
	CreateOrder(orderid, projectid string) error
	AddOrderItem(orderid, productid, name, uom string) error
	RemoveOrderItem(orderid, productid string) error
	ModifyRequestedQuantity(orderid, productid string, quantity int) error
	SendOrder(orderid string) error
	ReceiveOrderItem(orderid, productid string, quantity int) error
	FindOrder(orderid string) (*supply.Order, error)
	FindProjectOrderDates(projectid string) ([]ProjectOrder, error)
}

type orderRepository interface {
	supply.OrderRepository
	FindDates(projectid string) ([]ProjectOrder, error)
}

type service struct {
	order orderRepository
}

func NewOrderingService(order orderRepository) *service {
	return &service{order: order}
}

func (s *service) CreateOrder(orderid, projectid string) error {
	order := supply.Create(orderid, projectid)

	err := s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddOrderItem(orderid, productid, name, uom string) error {
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

func (s *service) RemoveOrderItem(orderid, productid string) error {
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

func (s *service) ModifyRequestedQuantity(orderid, productid string, quantity int) error {
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

func (s *service) SendOrder(orderid string) error {
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

func (s *service) ReceiveOrderItem(orderid, productid string, quantity int) error {
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

func (s *service) FindOrder(orderid string) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}

// Order read models

type ProjectOrder struct {
	OrderID  string
	SentDate int64
	Status   supply.OrderStatus
}

func (s *service) FindProjectOrderDates(projectid string) ([]ProjectOrder, error) {
	orders, err := s.order.FindDates(projectid)
	if err != nil {
		return []ProjectOrder{}, err
	}
	return orders, nil
}
