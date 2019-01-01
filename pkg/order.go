package material

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound    = errors.New("order not found")
	ErrOrderAlreadySent = errors.New("order already sent")
	ErrMustHaveItems    = errors.New("order must have at least 1 item")
	ErrQuantityZero     = errors.New("item quantity must be greater than 0")
)

type (
	OrderID   string
	ProjectID string
)

type Order struct {
	OrderID   OrderID
	ProjectID ProjectID
	version   int
	Statuses  []Status
	List      List
}

type Status struct {
	Date time.Time
	Type OrderStatus
}

func NewOrder(o OrderID, p ProjectID) (*Order, error) {
	return &Order{
		OrderID:   o,
		ProjectID: p,
		version:   0,
		Statuses: []Status{
			{
				Date: time.Now(),
				Type: New,
			},
		},
		List: List{
			Items: nil,
		},
	}, nil
}

func (o *Order) SendOrder() error {
	switch {
	case !o.List.haveItems():
		return ErrMustHaveItems
	case o.List.missingQuantities():
		return ErrQuantityZero
	case o.alreadySent():
		return ErrOrderAlreadySent
	default:
		o.send()
		return nil
	}
}

func (o *Order) AddItemToList(id ProductID, name string, uom UOM) error {
	return o.List.addItem(id, name, uom)
}

func (o *Order) RemoveItemFromList(id ProductID) error {
	return o.List.removeItem(id)
}

func (o *Order) UpdateQuantityRequested(id ProductID, q QuantityRequested) error {
	return o.List.updateQuantityRequested(id, q)
}

func (o *Order) ReceiveQuantity(id ProductID, q QuantityReceived) error {
	return o.List.receiveQuantity(id, q)
}

func (o *Order) send() {
	o.Statuses = append(o.Statuses, Status{
		Date: time.Time{},
		Type: Sent,
	})
}

func (o *Order) alreadySent() bool {
	return o.Statuses[len(o.Statuses)-1].Type != New
}

type OrderStatus int

const (
	New OrderStatus = iota
	Sent
	Received
	Processing
	OnRoute
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case New:
		return "New"
	case Sent:
		return "Sent"
	case Received:
		return "Received"
	case Processing:
		return "Processing"
	case OnRoute:
		return "OnRoute"
	case Complete:
		return "Complete"
	}
	return ""
}
