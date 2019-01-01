package material

import (
	"errors"
	"fmt"
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
	Statuses  []Status
	List      List
	OrderPOs  OrderPOs
	// Project string
	// version int
}

type Status struct {
	Date time.Time
	Type OrderStatus
}

func NewOrder(o OrderID, p ProjectID) (*Order, error) {
	return &Order{
		OrderID:   o,
		ProjectID: p,
		Statuses: []Status{
			{
				Date: time.Now(),
				Type: New,
			},
		},
		List: List{
			Items: nil,
		},
		OrderPOs: OrderPOs{
			POs: nil,
		},
	}, nil
}

func (o *Order) SendOrder() error {
	switch {
	case o.List.Items == nil:
		fmt.Println("nil")
		return ErrMustHaveItems
	case o.List.missingQuantities():
		return ErrQuantityZero
	case o.alreadySent():
		return ErrOrderAlreadySent
	default:
		o.updateStatus(Sent)
		return nil
	}
}

func (o *Order) updateStatus(s OrderStatus) {
	o.Statuses = append(o.Statuses, Status{
		Date: time.Time{},
		Type: s,
	})
}

func (o *Order) alreadySent() bool {
	return o.Statuses[len(o.Statuses)-1].Type != New
}

func (o *Order) ReceiveQuantity(id ProductID, q QuantityReceived) error {
	if err := o.List.receiveQuantity(id, q); err != nil {
		return err
	}
	if o.List.receivedAll() {
		o.updateStatus(Complete)
	}
	return nil
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

func (o *Order) AddPOtoOrder(n PONumber, s Supplier) error {
	return o.OrderPOs.add(n, s)
}

func (o *Order) RemovePOfromOrder(n PONumber) error {
	return o.OrderPOs.remove(n)
}

func (o *Order) UpdateItemPO(id ProductID, n PONumber, s Supplier) error {
	return o.List.updateItemPO(id, n, s)
}

func (o *Order) RemoveItemPO(id ProductID, n PONumber) error {
	return o.List.removePOfromItem(id, n)
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
