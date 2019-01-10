package purchasing

import "supply/pkg"

type PurchasingService interface {
	UpdateProduct(name, uom string)
	AddProduct(id, name, uom string) *supply.Product
	UpdateItemPO(orderid, productid, ponumber string) error
	Process(orderid string)
}

type Service struct {
	product supply.ProductRepository
	order   supply.OrderRepository
}

func NewPurchasingService(p supply.ProductRepository, o supply.OrderRepository) *Service {
	return &Service{
		product: p,
		order:   o,
	}
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
