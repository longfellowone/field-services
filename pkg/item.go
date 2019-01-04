package field

type ProductUUID string

type Item struct {
	ProductUUID
	Name string
	UOM
	QuantityRequested uint
	QuantityReceived  uint
	ItemStatus
	PurchaseOrder
	Index int
}

func newItem(id ProductUUID) Item {
	return Item{
		ProductUUID:       id,
		QuantityRequested: 0,
		QuantityReceived:  0,
	}
}

func (i Item) receiveItem(quantity uint) Item {
	i.QuantityReceived += quantity

	switch {
	case i.QuantityReceived >= i.QuantityRequested:
		i.ItemStatus = Filled
	case quantity > 0 && quantity < i.QuantityRequested:
		i.ItemStatus = BackOrdered
	}
	return i
}

func (i Item) updateQuantityRequested(quantity uint) Item {
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
