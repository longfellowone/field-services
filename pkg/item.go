package material

import (
	"time"
)

type (
	ProductID         string
	QuantityRequested int
	QuantityReceived  int
)

type Item struct {
	ProductID         ProductID
	Name              string
	UOM               UOM
	QuantityRequested QuantityRequested
	QuantityReceived  QuantityReceived
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

func (l Item) receive(q QuantityReceived) Item {
	rec := int(q)
	req := int(l.QuantityRequested)

	l.LastUpdate = time.Now()
	l.QuantityReceived = +q

	switch {
	case rec >= req:
		l.Status = Filled
	case rec < req:
		l.Status = BackOrdered
	default:
		l.Status = Waiting
	}
	return l
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
