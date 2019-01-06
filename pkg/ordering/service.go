package ordering

import (
	"github.com/pkg/errors"
	"log"
	"supply/pkg"
)

type Service struct {
	db supply.OrderRepository
}

func NewOrderingService(db supply.OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateOrder(orderid supply.OrderUUID, projectid supply.ProjectUUID) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}

	if order.OrderUUID == orderid {
		return errors.New("order already exists")
	}
	order = supply.Create(orderid, projectid)

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendOrder(orderid supply.OrderUUID) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}
	order.Send()

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddItemToOrder(orderid supply.OrderUUID, productid supply.ProductUUID, name string, uom supply.UOM) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}
	order.AddItem(productid, name, uom)

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveItemFromOrder(orderid supply.OrderUUID, productid supply.ProductUUID) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}
	order.RemoveItem(productid)

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReceiveOrderItem(orderid supply.OrderUUID, productid supply.ProductUUID, quantity uint) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}
	order.ReceiveItem(productid, quantity)

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ModifyRequestedQuantity(orderid supply.OrderUUID, productid supply.ProductUUID, quantity uint) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}
	order.UpdateQuantityRequested(productid, quantity)

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateItemPO(orderid supply.OrderUUID, productid supply.ProductUUID, ponumber string) error {
	order, err := s.db.Find(orderid)
	if err != nil {
		return err
	}
	order.UpdatePO(productid, ponumber)

	err = s.db.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FindAllProjectOrders(uuid supply.ProjectUUID) ([]supply.Order, error) {
	findAll, err := s.db.FindAllFromProject(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}

func (s *Service) FindOrder(uuid supply.OrderUUID) (*supply.Order, error) {
	findAll, err := s.db.Find(uuid)
	if err != nil {
		log.Println(err)
	}
	return findAll, nil
}
