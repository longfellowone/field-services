package material

import (
	"errors"
	"time"
)

var (
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
	ErrItemQuantityZero  = errors.New("item quantity must be greater than 0")
)

type ProductID string

type List struct {
	Items []Item
}

type Item struct {
	ProductID         ProductID
	Name              string
	UOM               UOM
	QuantityRequested QuantityRequested
	QuantityReceived  QuantityReceived
	Status            ItemStatus
	LastUpdate        time.Time
	PO                string
}

type QuantityRequested int
type QuantityReceived int

func (l *List) addItem(id ProductID, name string, uom UOM) error {
	if l.itemExists(id) {
		return ErrItemAlreadyOnList
	}
	l.Items = append(l.Items, newItem(id, name, uom))
	return nil
}

func (l *List) removeItem(id ProductID) error {
	for i := range l.Items {
		if l.Items[i].ProductID == id {
			l.remove(i, &l.Items)
			return nil
		}
	}
	return ErrItemNotFound
}

func (l *List) updateQuantityRequested(id ProductID, q QuantityRequested) error {
	if isZero(int(q)) {
		return ErrItemQuantityZero
	}

	for i, item := range l.Items {
		if item.ProductID == id {
			l.Items[i].QuantityRequested = q
			return nil
		}
	}
	return ErrItemNotFound
}

func (l *List) receiveQuantity(id ProductID, q QuantityReceived) error {
	if isZero(int(q)) {
		return ErrItemQuantityZero
	}

	for i, item := range l.Items {
		if item.ProductID == id {
			l.Items[i].receive(q)
			return nil
		}
	}
	return ErrItemNotFound
}

func (l *List) itemExists(id ProductID) bool {
	for _, item := range l.Items {
		if item.ProductID == id {
			return true
		}
	}
	return false
}

func (l *List) missingQuantities() bool {
	for _, item := range l.Items {
		if item.QuantityRequested <= 0 {
			return true
		}
	}
	return false
}

func (l *List) haveItems() bool {
	return len(l.Items) != 0
}

func (l *Item) receive(q QuantityReceived) {
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
}

func (l *List) remove(i int, item *[]Item) *[]Item {
	copy(l.Items[i:], l.Items[i+1:])
	l.Items[len(l.Items)-1] = Item{}
	l.Items = l.Items[:len(l.Items)-1]
	return item
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
		PO:                "",
	}
}

func isZero(q int) bool {
	return q <= 0
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
