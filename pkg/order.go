package supply

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrMustHaveItems     = errors.New("order must have at least 1 Item")
	ErrQuantityZero      = errors.New("quantity of all order Items must be greater than 0")
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
)

type OrderRepository interface {
	Save(o *Order) error
	Find(id string) (*Order, error)
	FindAllFromProject(id string) ([]Order, error)
}

type Order struct {
	OrderID   string
	ProjectID string
	Items     []Item
	OrderDate time.Time
	Status    OrderStatus
}

// Returns a new *Order
func Create(id, pid string) *Order {
	return &Order{
		OrderID:   id,
		ProjectID: pid,
		Items:     []Item{},
		Status:    New,
	}
}

// Adds an item to order
func (o *Order) AddItem(id, name, uom string) error {
	_, err := o.findItem(id)
	if err == nil {
		return ErrItemAlreadyOnList
	}
	o.Items = append(o.Items, *newItem(id, name, uom))
	return nil
}

// Removes an item from order
func (o *Order) RemoveItem(id string) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.Items = append(o.Items[:i], o.Items[i+1:]...)
	return nil
}

// Updates the quantity requested of a single order item
func (o *Order) UpdateQuantityRequested(id string, quantity uint) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.Items[i].QuantityRequested = quantity
	return nil
}

// Marks an order sent
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

// Updates the PO number an item
func (o *Order) UpdatePO(id, po string) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	if po == "" {
		po = "N/A"
	}
	o.Items[i].PONumber = po
	return nil
}

// Updates the quantity received of an item
func (o *Order) ReceiveItem(id string, quantity uint) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}

	o.Items[i].receive(quantity)

	if o.receivedAll() {
		o.Status = Complete
	}
	return nil
}

// Finds index of an item
func (o *Order) findItem(id string) (int, error) {
	for i := range o.Items {
		if o.Items[i].ProductID == id {
			return i, nil
		}
	}
	return 0, ErrItemNotFound
}

// Checks to make sure all items have been received
func (o *Order) receivedAll() bool {
	for i := range o.Items {
		if o.Items[i].ItemStatus == Waiting || o.Items[i].ItemStatus == BackOrdered {
			return false
		}
	}
	return true
}

// Checks an order to make sure all items requested have a quantity
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
	default:
		return fmt.Sprintf("%d", int(s))
	}
}
