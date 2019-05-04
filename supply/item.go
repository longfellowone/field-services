package supply

import (
	"fmt"
	"time"
)

type Item struct {
	ID                string
	Name              string
	UOM               string
	QuantityRequested int
	QuantityReceived  int
	QuantityRemaining int
	ItemStatus        ItemStatus
	PONumber          string
	DateAdded         int64
	Removed           bool
}

// Returns a new *Item
func newItem(id, name, uom string) *Item {
	return &Item{
		ID:         id,
		Name:       name,
		UOM:        uom,
		ItemStatus: New,
		PONumber:   "N/A",
		DateAdded:  time.Now().Unix(),
	}
}

// Updates the quantity requested of an item
func (i *Item) updateRequested(quantity int) {
	i.QuantityRequested = quantity
}

// Updates quantity received, remaining and item status
func (i *Item) receive(quantity int) {
	i.QuantityReceived = quantity

	switch {
	case i.QuantityReceived == i.QuantityRequested:
		i.ItemStatus = Filled
		i.QuantityRemaining = 0

	case quantity > 0 && quantity < i.QuantityRequested:
		i.ItemStatus = BackOrdered
		i.QuantityRemaining = i.QuantityRequested - i.QuantityReceived

	case i.QuantityReceived > i.QuantityRequested:
		i.ItemStatus = OrderExceeded
		i.QuantityRemaining = 0
	}
}

type ItemStatus int

const (
	New ItemStatus = iota
	Waiting
	Filled
	BackOrdered
	OrderExceeded
	NotOrdered
)

func (s ItemStatus) String() string {
	switch s {
	case New:
		return "New"
	case Waiting:
		return "Waiting"
	case Filled:
		return "Filled"
	case BackOrdered:
		return "Back Ordered"
	case OrderExceeded:
		return "Order Exceeded"
	case NotOrdered:
		return "Not Ordered"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}
