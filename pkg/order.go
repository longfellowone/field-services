package material

import (
	"errors"
	"time"
)

type OrderID string
type ProjectID string

type Order struct {
	OrderID   OrderID
	Project   ProjectID
	Date      time.Time
	Status    OrderStatus
	LineItems []LineItem
}

func NewOrder(o OrderID, p ProjectID) (*Order, error) {
	return &Order{
		OrderID: o,
		Project: p,
		Date:    time.Now(),
		Status:  BackOrdered,
	}, nil
}

type LineItem struct {
}

var ErrMyError = errors.New("message")

type OrderStatus int

const (
	Complete OrderStatus = iota
	BackOrdered
)

func (s OrderStatus) String() string {
	switch s {
	case Complete:
		return "Complete"
	case BackOrdered:
		return "Back Ordered"
	}
	return ""
}
