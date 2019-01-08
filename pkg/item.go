package supply

type Item struct {
	ProductID         string
	Name              string
	UOM               string
	QuantityRequested uint
	QuantityReceived  uint
	QuantityRemaining uint
	ItemStatus
	PONumber string
}

// Returns a new *Item
func newItem(id, name, uom string) *Item {
	return &Item{
		ProductID:  id,
		Name:       name,
		UOM:        uom,
		ItemStatus: Waiting,
		PONumber:   "N/A",
	}
}

// Updates the quantity requested of an item
func (i *Item) updateRequested(quantity uint) {
	i.QuantityRequested = quantity
}

// Updates quantity received, remaining and item status
func (i *Item) receive(quantity uint) {
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
	Waiting ItemStatus = iota
	Filled
	BackOrdered
	OrderExceeded
	NotOrdered
)

func (s ItemStatus) String() string {
	switch s {
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
	}

	return ""
}
