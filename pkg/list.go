package material

import (
	"errors"
	"time"
)

var (
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
	ErrQuantityZero      = errors.New("item quantity must be greater than 0")
)

type ProductID string

type List struct {
	Items []Item
}

type QuantityRequested int
type QuantityReceived int

func (l *List) addItem(id ProductID, name string, uom UOM) error {
	if err := l.itemAlreadyExists(id); err != nil {
		return ErrItemAlreadyOnList
	}

	l.Items = append(l.Items, Item{
		ProductID:         id,
		Name:              name,
		UOM:               uom,
		QuantityRequested: 0,
		QuantityReceived:  0,
		Status:            Waiting,
		LastUpdate:        time.Now(),
		PO:                "",
	})
	return nil
}

func (l *List) removeItem(id ProductID) error {
	for i := range l.Items {
		if l.Items[i].ProductID == id {
			copy(l.Items[i:], l.Items[i+1:])
			l.Items[len(l.Items)-1] = Item{}
			l.Items = l.Items[:len(l.Items)-1]
			return nil
		}
	}
	return ErrItemNotFound
}

func (l *List) updateQuantityRequested(id ProductID, q QuantityRequested) error {
	if q <= 0 {
		return ErrQuantityZero
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
	if q <= 0 {
		return ErrQuantityZero
	}

	for i, item := range l.Items {
		if item.ProductID == id {
			l.Items[i].receive(q)
			return nil
		}
	}
	return ErrItemNotFound
}

func (l *List) itemAlreadyExists(id ProductID) error {
	for _, item := range l.Items {
		if item.ProductID == id {
			return ErrItemAlreadyOnList
		}
	}
	return nil
}

func (l *List) haveItems() bool {
	if len(l.Items) <= 0 {
		return false
	}
	return true
}

func (l *List) missingQuantities() error {
	for _, item := range l.Items {
		if item.QuantityRequested <= 0 {
			return ErrQuantityZero
		}
	}
	return nil
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

func (l *Item) receive(q QuantityReceived) {
	l.LastUpdate = time.Now()
	l.QuantityReceived = +q
	switch {
	case int(q) >= int(l.QuantityRequested):
		l.Status = Filled
	case int(q) < int(l.QuantityRequested):
		l.Status = BackOrdered
	default:
		l.Status = Waiting
	}
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
