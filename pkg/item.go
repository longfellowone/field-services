package supply

type Item struct {
	ProductUUID
	Name              string
	UOM               UOM
	QuantityRequested uint
	QuantityReceived  uint
	QuantityRemaining uint
	ItemStatus
	PONumber string
}

func newItem(uuid ProductUUID, name string, uom UOM) Item {
	return Item{
		ProductUUID: uuid,
		Name:        name,
		UOM:         uom,
		ItemStatus:  Waiting,
		PONumber:    "N/A",
	}
}

func (i Item) receive(quantity uint) Item {
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
	return i
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
