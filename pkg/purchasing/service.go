package purchasing

import (
	"supply/pkg"
)

type PurchasingService interface {
	UpdateItemPO(orderid, productid, ponumber string) error
	Process(orderid string) error
	UpdateProduct(id, category, name, uom string) error
	AddProduct(id, category, name, uom string) error
}

type ProductRepository interface {
	supply.ProductRepository
}

type OrderRepository interface {
	supply.OrderRepository
}

type Service struct {
	product ProductRepository
	order   OrderRepository
}

func NewPurchasingService(product ProductRepository, order OrderRepository) (*Service, error) {
	return &Service{
		product: product,
		order:   order,
	}, nil
}

func (s *Service) UpdateItemPO(orderid, productid, ponumber string) error {
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

func (s *Service) Process(orderid string) error {
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

func (s *Service) UpdateProduct(id, category, name, uom string) error {
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

func (s *Service) AddProduct(id, category, name, uom string) error {
	product := supply.NewProduct(id, category, name, uom)

	err := s.product.Save(product)
	if err != nil {
		return err
	}
	return nil
}
