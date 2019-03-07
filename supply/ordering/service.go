package ordering

import (
	"field/supply"
	"log"
)

type Service interface {
	CreateOrder(orderid, projectid string) error
	AddOrderItem(orderid, productid, name, uom string) error
	RemoveOrderItem(orderid, productid string) error
	ModifyRequestedQuantity(orderid, productid string, quantity int) error
	SendOrder(orderid string) error
	ReceiveOrderItem(orderid, productid string, quantity int) error
	FindOrder(orderid string) *supply.Order
	FindProjectOrderDates(projectid string) []ProjectOrder
}

type orderRepository interface {
	supply.OrderRepository
	FindDates(projectid string) ([]ProjectOrder, error)
}

type Svc struct {
	order orderRepository
}

func NewOrderingService(order orderRepository) *Svc {
	return &Svc{order: order}
}

func (s *Svc) CreateOrder(orderid, projectid string) error {
	order := supply.Create(orderid, projectid)

	err := s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) AddOrderItem(orderid, productid, name, uom string) error {
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

func (s *Svc) RemoveOrderItem(orderid, productid string) error {
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

func (s *Svc) ModifyRequestedQuantity(orderid, productid string, quantity int) error {
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

func (s *Svc) SendOrder(orderid string) error {
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

func (s *Svc) ReceiveOrderItem(orderid, productid string, quantity int) error {
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

func (s *Svc) FindOrder(orderid string) *supply.Order {
	order, err := s.order.Find(orderid)
	if err != nil {
		log.Println(err)
		return &supply.Order{}
	}
	return order
}

// Order read models

type ProjectOrder struct {
	OrderID  string
	SentDate int64
	Status   supply.OrderStatus
}

func (s *Svc) FindProjectOrderDates(projectid string) []ProjectOrder {
	orders, err := s.order.FindDates(projectid)
	if err != nil {
		log.Println(err)
		return []ProjectOrder{}
	}
	return orders
}
