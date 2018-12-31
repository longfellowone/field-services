package material

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type OrderID string
type ProjectID string

type Order struct {
	OrderID   OrderID
	ProjectID ProjectID
	version   int
	statuses  []status
	lineItems []lineItem
}

func NewOrder(o OrderID, p ProjectID) (*Order, error) {
	return &Order{
		OrderID:   o,
		ProjectID: p,
		version:   0,
		statuses: []status{
			{
				Date: time.Now(),
				Type: NotSent,
			},
		},
		lineItems: nil,
	}, nil
}

type status struct {
	Date time.Time
	Type OrderStatus
}

type lineItem struct {
	ProductID string
	Name      string
	UOM       UOM
	Quantity  int
	Status    LineItemStatus
	PO        string
}

var ErrMyError = errors.New("message")

type OrderStatus int

const (
	NotSent OrderStatus = iota
	Sent
	Received
	Processing
	OnRoute
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case NotSent:
		return "Not Sent"
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

type LineItemStatus int

const (
	Waiting LineItemStatus = iota
	Filled
	BackOrdered
)

func (s LineItemStatus) String() string {
	switch s {
	case Waiting:
		return "Waiting"
	case Filled:
		return "Filled"
	case BackOrdered:
		return "Back Ordered"
	}
	return ""
}

type UOM int

const (
	EA UOM = iota
	FT
	M
)

func (s UOM) String() string {
	switch s {
	case EA:
		return "ea"
	case FT:
		return "ft"
	case M:
		return "m"
	}
	return ""
}
