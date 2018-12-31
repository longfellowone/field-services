package material

type ProductID string

type LineItem struct {
	ProductID ProductID
	Name      string
	UOM       UOM
	Quantity  int
	status    LineItemStatus
	PO        string
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
