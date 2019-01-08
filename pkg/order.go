package supply

import (
	"errors"
	"time"
)

var (
	ErrMustHaveItems     = errors.New("order must have at least 1 item")
	ErrQuantityZero      = errors.New("quantity of all order Items must be greater than 0")
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
)

type OrderRepository interface {
	Save(o *Order) error
	Find(uuid string) (*Order, error)
	FindAllFromProject(uuid string) ([]Order, error)
}

type Order struct {
	OrderUUID   string
	ProjectUUID string
	Items       []item
	OrderDate   time.Time
	Status      OrderStatus
}

func Create(id, pid string) *Order {
	return &Order{
		OrderUUID:   id,
		ProjectUUID: pid,
		Status:      New,
	}
}

func (o *Order) Send() error {
	switch {
	case len(o.Items) == 0:
		return ErrMustHaveItems
	case o.missingQuantities():
		return ErrQuantityZero
	}

	o.Status = Sent
	o.OrderDate = time.Now()
	return nil
}

func (o *Order) AddItem(uuid, name, uom string) error {
	_, err := o.findItem(uuid)
	if err == nil {
		return ErrItemAlreadyOnList
	}
	o.Items = append(o.Items, newItem(uuid, name, uom))
	return nil
	// Test for len() then test [len()-1}
}

func (o *Order) RemoveItem(id string) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.Items = append(o.Items[:i], o.Items[i+1:]...)
	return nil
}

func (o *Order) UpdateQuantityRequested(uuid string, quantity uint) error {
	i, err := o.findItem(uuid)
	if err != nil {
		return err
	}
	o.Items[i].QuantityRequested = quantity
	return nil
}

func (o *Order) ReceiveItem(uuid string, quantity uint) error {
	i, err := o.findItem(uuid)
	if err != nil {
		return err
	}

	o.Items[i].receive(quantity)

	if o.receivedAll() {
		o.Status = Complete
	}
	return nil
}

func (o *Order) UpdatePO(uuid, po string) error {
	i, err := o.findItem(uuid)
	if err != nil {
		return err
	}
	if po == "" {
		po = "N/A"
	}
	o.Items[i].PONumber = po
	return nil
}

func (o *Order) findItem(uuid string) (int, error) {
	for i := range o.Items {
		if o.Items[i].ProductUUID == uuid {
			return i, nil
		}
	}
	return 0, ErrItemNotFound
}

func (o *Order) receivedAll() bool {
	for i := range o.Items {
		if o.Items[i].ItemStatus == Waiting {
			return false
		}
	}
	return true
}

func (o *Order) missingQuantities() bool {
	for i := range o.Items {
		if o.Items[i].QuantityRequested == 0 {
			return true
		}
	}
	return false
}

type OrderStatus int

const (
	New OrderStatus = iota
	Sent
	OnRoute
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case New:
		return "New"
	case Sent:
		return "Sent"
	case OnRoute:
		return "OnRoute"
	case Complete:
		return "Complete"
	}
	return ""
}
