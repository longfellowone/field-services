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

func newItem(id ProductUUID, quantity int) Item {
	return Item{
		ProductUUID:       id,
		QuantityRequested: quantity,
		QuantityReceived:  0,
	}
}

func (i Item) ReceiveItem(quantity int) Item {
	i.QuantityReceived = quantity
	return i
}

func (i Item) UpdateQuantityRequested(quantity int) Item {
	i.QuantityRequested = quantity
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
