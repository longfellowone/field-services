package material

import (
	"errors"
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

type Item struct {
	ProductID         ProductID
	Name              string
	UOM               UOM
	QuantityRequested QuantityRequested
	QuantityReceived  QuantityReceived
	status            LineItemStatus
	PO                string
}

type QuantityRequested int
type QuantityReceived int

func (l *List) addItem(id ProductID, name string, uom UOM) error {
	if err := l.itemAlreadyExists(id); err != nil {
		return err
	}

	l.Items = append(l.Items, Item{
		ProductID:         id,
		Name:              name,
		UOM:               uom,
		QuantityRequested: 0,
		QuantityReceived:  0,
		status:            Waiting,
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

func (l *List) updateQuantityReceived(id ProductID, q QuantityReceived) error {
	return nil
}

func (l *List) itemAlreadyExists(id ProductID) error {
	for _, item := range l.Items {
		if item.ProductID == id {
			return ErrItemAlreadyOnList
		}
	}
	return nil
}

func (l *List) hasItems() error {
	if len(l.Items) <= 0 {
		return ErrMustHaveItems
	}
	return nil
}

func (l *List) missingQuantities() error {
	for _, item := range l.Items {
		if item.QuantityRequested <= 0 {
			return ErrQuantityZero
		}
	}
	return nil
}

//func (i Item) itemMissingQuantity() error {
//	if i.QuantityRequested <= 0 {
//		return ErrQuantityZero
//	}
//	return nil
//}

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
