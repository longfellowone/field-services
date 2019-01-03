package material

import (
	"time"
)

type ProductID string

type Item struct {
	ProductID         ProductID
	Name              string
	UOM               UOM
	QuantityRequested int
	QuantityReceived  int
	Status            ItemStatus
	LastUpdate        time.Time
	PO                PurchaseOrder
}

func newItem(id ProductID, name string, uom UOM) Item {
	return Item{
		ProductID:         id,
		Name:              name,
		UOM:               uom,
		QuantityRequested: 0,
		QuantityReceived:  0,
		Status:            Waiting,
		LastUpdate:        time.Now(),
		PO:                PurchaseOrder{},
	}
}

func (i *Item) adjustQuantity(quantity int) {
	i.QuantityRequested = quantity
}

func (i *Item) receive(received int) {
	requested := i.QuantityRequested

	i.LastUpdate = time.Now()
	i.QuantityReceived = +received

	switch {
	case received >= requested:
		i.Status = Filled
	case received < requested:
		i.Status = BackOrdered
	default:
		i.Status = Waiting
	}
}

func (i *Item) updatePO(number string, supplier string) {
	i.PO = newPO(number, supplier)
}

func (i *Item) removePO() {
	i.PO = PurchaseOrder{}
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
