package material

import (
	"errors"
)

var (
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
	ErrItemQuantityZero  = errors.New("item quantity must be greater than 0")
)

type List struct {
	Items []Item
}

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
			l.remove(i)
			return nil
		}
	}
	return ErrItemNotFound
}

func (l *List) remove(i int) {
	copy(l.Items[i:], l.Items[i+1:])
	l.Items[len(l.Items)-1] = Item{}
	l.Items = l.Items[:len(l.Items)-1]
}

func (l *List) updateQuantityRequested(id ProductID, q QuantityRequested) error {
	if q <= 0 {
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
	if q <= 0 {
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

func (l *List) receivedAll() bool {
	for _, item := range l.Items {
		if item.Status != Filled {
			return false
		}
	}
	return true
}

func (l *List) updateItemPO(id ProductID, n PONumber, s Supplier) error {
	for i := range l.Items {
		item := &l.Items[i]
		if item.ProductID == id {
			item.updatePO(n, s)
			return nil
		}
	}
	return ErrPOnotFound
}

func (l *List) removePOfromItem(id ProductID, n PONumber) error {
	return nil
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
