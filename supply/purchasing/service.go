package purchasing

import (
	"field/supply"
)

type Service interface {
	UpdateItemPO(orderid, productid, ponumber string) error
	Process(orderid string) error
	UpdateProduct(id, category, name, uom string) error
	AddProduct(id, category, name, uom string) error
}

type productRepository interface {
	supply.ProductRepository
}

type orderRepository interface {
	supply.OrderRepository
}

type Svc struct {
	product productRepository
	order   orderRepository
}

func NewPurchasingService(product productRepository, order orderRepository) *Svc {
	return &Svc{
		product: product,
		order:   order,
	}
}

func (s *Svc) UpdateItemPO(orderid, productid, ponumber string) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	err = order.UpdatePO(productid, ponumber)
	if err != nil {
		return err
	}

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) Process(orderid string) error {
	order, err := s.order.Find(orderid)
	if err != nil {
		return err
	}

	order.Process()

	err = s.order.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) UpdateProduct(id, category, name, uom string) error {
	product, err := s.product.Find(id)
	if err != nil {
		return err
	}

	product.ModifyProduct(category, name, uom)

	err = s.product.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) AddProduct(id, category, name, uom string) error {
	product := supply.NewProduct(id, category, name, uom)

	err := s.product.Save(product)
	if err != nil {
		return err
	}
	return nil
}
