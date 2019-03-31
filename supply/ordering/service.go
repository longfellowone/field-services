package ordering

import (
	"field/supply"
)

type Service interface {
	// Orders
	CreateOrder(orderid, projectid, name, foreman, email string) (*supply.Order, error)
	AddOrderItem(orderid, productid, name, uom string) (*supply.Order, error)
	RemoveOrderItem(orderid, productid string) (*supply.Order, error)
	ModifyRequestedQuantity(orderid, productid string, quantity int) (*supply.Order, error)
	SendOrder(orderid string) (*supply.Order, error)
	ReceiveOrderItem(orderid, productid string, quantity int) (*supply.Order, error)
	FindOrder(orderid string) (*supply.Order, error)
	FindProjectOrderDates(projectid string) ([]ProjectOrder, error)
	// Projects
	CreateProject(projectid, name, foreman, email string) (*supply.Project, error)
	CloseProject(projectid string) (*supply.Project, error)
	FindProjectsByForeman(foremanid string) ([]supply.Project, error)
}

type orderRepository interface {
	supply.OrderRepository
	FindDates(projectid string) ([]ProjectOrder, error)
}

type projectRepository interface {
	supply.ProjectRepository
	FindAllByForeman(foremanid string) ([]supply.Project, error)
}

type service struct {
	order   orderRepository
	project projectRepository
}

func NewOrderingService(order orderRepository, project projectRepository) *service {
	return &service{order: order, project: project}
}

func (s *service) CreateOrder(orderid, projectid, name, foreman, email string) (*supply.Order, error) {
	order := supply.Create(orderid, projectid, name, foreman, email)

	err := s.order.Save(order)
	if err != nil {
		return &supply.Order{}, nil
	}
	return order, nil
}

func (s *service) AddOrderItem(orderid, productid, name, uom string) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = order.AddItem(productid, name, uom)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = s.order.Save(order)
	if err != nil {
		return &supply.Order{}, nil
	}
	return order, nil
}

func (s *service) RemoveOrderItem(orderid, productid string) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = order.RemoveItem(productid)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = s.order.Save(order)
	if err != nil {
		return &supply.Order{}, nil
	}
	return order, nil
}

func (s *service) ModifyRequestedQuantity(orderid, productid string, quantity int) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = order.UpdateQuantityRequested(productid, quantity)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = s.order.Save(order)
	if err != nil {
		return &supply.Order{}, nil
	}
	return order, nil
}

func (s *service) SendOrder(orderid string) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = order.Send()
	if err != nil {
		return &supply.Order{}, nil
	}

	err = s.order.Save(order)
	if err != nil {
		return &supply.Order{}, nil
	}
	return order, nil
}

func (s *service) ReceiveOrderItem(orderid, productid string, quantity int) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = order.ReceiveItem(productid, quantity)
	if err != nil {
		return &supply.Order{}, nil
	}

	err = s.order.Save(order)
	if err != nil {
		return &supply.Order{}, nil
	}
	return order, nil
}

func (s *service) FindOrder(orderid string) (*supply.Order, error) {
	order, err := s.order.Find(orderid)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}

func (s *service) CreateProject(projectid, name, foreman, email string) (*supply.Project, error) {
	project := supply.NewProject(projectid, name, foreman, email)
	err := s.project.Save(project)
	if err != nil {
		return &supply.Project{}, nil
	}
	return project, nil
}
func (s *service) CloseProject(projectid string) (*supply.Project, error) {
	project, err := s.project.Find(projectid)
	if err != nil {
		return &supply.Project{}, nil
	}

	project.Close()

	err = s.project.Save(project)
	if err != nil {
		return &supply.Project{}, nil
	}
	return project, nil
}
func (s *service) FindProjectsByForeman(foremanid string) ([]supply.Project, error) {
	projects, err := s.project.FindAllByForeman(foremanid)
	if err != nil {
		return []supply.Project{}, nil
	}
	return projects, nil
}

// Order read models

type ProjectOrder struct {
	ID       string
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
