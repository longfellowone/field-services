package field

type ProductUUID string

type Item struct {
	ProductUUID
	Name string
	UOM
	QuantityRequested int
	QuantityReceived  int
	ItemStatus
	PurchaseOrder
}

func newItem(id ProductUUID) Item {
	return Item{
		ProductUUID:       id,
		QuantityRequested: 0,
		QuantityReceived:  0,
	}
}

func (i Item) receiveItem(quantity int) Item {
	i.QuantityReceived += quantity
	if i.QuantityReceived >= i.QuantityRequested {
		i.updateStatus(Filled)
	}
	return i
}

func (i Item) updateQuantityRequested(quantity int) Item {
	i.QuantityRequested = quantity
	return i
}

func (i Item) updateStatus(status ItemStatus) Item {
	i.ItemStatus = status
	return i
}

type ItemStatus int

const (
	Waiting ItemStatus = iota
	Filled
	BackOrdered
)

func (s ItemStatus) String() string {
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
