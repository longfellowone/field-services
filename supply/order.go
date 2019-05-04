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
	Delete(id string) error
}

type Order struct {
	ID string
	Project
	Items    []*Item
	SentDate int64
	Status   OrderStatus
	Comments string
}

// Returns a new *Order
func Create(id, pid, name, foreman, email string) *Order {
	return &Order{
		ID: id,
		Project: Project{
			ID:           pid,
			Name:         name,
			Foreman:      foreman,
			ForemanEmail: email,
		},
		Items:    []*Item{},
		SentDate: time.Now().Unix(),
		Status:   Draft,
	}
}

// Adds an item to order
func (o *Order) AddItem(id, name, uom string) error {
	_, err := o.findItem(id)
	if err == nil {
		return ErrItemAlreadyOnList
	}
	item := newItem(id, name, uom)
	o.Items = append([]*Item{item}, o.Items...)
	return nil
}

// Removes an item from order
func (o *Order) RemoveItem(id string) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.Items[i].Removed = true
	return nil
}

// Updates the quantity requested of a single order item
func (o *Order) UpdateQuantityRequested(id string, quantity int) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.Items[i].QuantityRequested = quantity
	return nil
}

// Marks an order sent
func (o *Order) Send(comments string) error {
	switch {
	case len(o.Items) == 0:
		return ErrMustHaveItems

	case o.missingQuantities():
		return ErrQuantityZero
	}

	for _, item := range o.Items {
		item.ItemStatus = Waiting
	}

	o.Comments = comments
	o.Status = Sent
	o.SentDate = time.Now().Unix()

	return nil
}

// Mark the order processed
func (o *Order) Process() {
	o.Status = Processed
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
func (o *Order) ReceiveItem(id string, quantity int) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}

	o.Items[i].receive(quantity)

	if o.receivedAll() {
		o.Status = Complete
		return nil
	}

	o.Status = Processed
	return nil
}

// Finds index of an item
func (o *Order) findItem(id string) (int, error) {
	for i := range o.Items {
		if o.Items[i].ID == id {
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
	Draft OrderStatus = iota
	Sent
	Processed
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case Draft:
		return "Draft"
	case Sent:
		return "Sent"
	case Processed:
		return "Processed"
	case Complete:
		return "Complete"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}
