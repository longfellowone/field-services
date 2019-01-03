package material

import (
	"errors"
	"log"
	"time"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrOrderAlreadySent  = errors.New("order already sent")
	ErrMustHaveItems     = errors.New("order must have at least 1 item")
	ErrQuantityZero      = errors.New("item quantity must be greater than 0")
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
	ErrItemQuantityZero  = errors.New("item quantity must be greater than 0")
	ErrPOalreadyExists   = errors.New("PO already exists")
	ErrPOnotFound        = errors.New("PO not found")
)

type OrderRepository interface {
	Save(o *Order) error
	Find(id OrderID) (*Order, error)
	FindAllFromProject(id ProjectID) ([]*Order, error)
}

type OrderID string
type ProjectID string

type Order struct {
	OrderID   OrderID
	ProjectID ProjectID
	Statuses  []Status
	List      []Item
	POs       []PurchaseOrder
}

func NewOrder(o OrderID, p ProjectID) *Order {
	return &Order{
		OrderID:   o,
		ProjectID: p,
		Statuses: []Status{
			{
				Date: time.Now(),
				Type: New,
			},
		},
		List: []Item{},
		POs:  []PurchaseOrder{},
	}
}

func (o *Order) AddItem(id ProductID, name string, uom UOM) {
	_, err := o.findItem(id)
	if err == nil {
		log.Println(ErrItemAlreadyOnList)
		return
	}
	item := newItem(id, name, uom)
	o.List = append(o.List, item)
}

func (o *Order) RemoveItem(id ProductID) {
	i, err := o.findItem(id)
	if err != nil {
		log.Println(ErrItemNotFound)
		return
	}
	o.List = append(o.List[:i], o.List[i+1:]...)
}

func (o *Order) UpdateQuantityRequested(id ProductID, quantity int) {
	if quantity <= 0 {
		log.Println(ErrItemQuantityZero)
		return
	}
	i, err := o.findItem(id)
	if err != nil {
		log.Println(ErrItemNotFound)
		return
	}
	o.List[i].adjustQuantity(quantity)
}

func (o *Order) ReceiveQuantity(id ProductID, quantity int) {
	if quantity <= 0 {
		log.Println(ErrItemQuantityZero)
		return
	}

	i, err := o.findItem(id)
	if err != nil {
		log.Println(ErrItemNotFound)
		return
	}

	o.List[i].receive(quantity)

	if o.receivedAll() {
		o.updateStatus(Complete)
	}
}

func (o *Order) SendOrder() {
	switch {
	case o.List == nil:
		log.Println(ErrMustHaveItems)
	case o.missingQuantities():
		log.Println(ErrQuantityZero)
	case o.alreadySent():
		log.Println(ErrOrderAlreadySent)
	default:
		o.updateStatus(Sent)
	}
}

func (o *Order) findItem(id ProductID) (int, error) {
	for i, item := range o.List {
		if item.ProductID == id {
			return i, nil
		}
	}
	return 0, ErrItemNotFound
}

func (o *Order) missingQuantities() bool {
	for _, item := range o.List {
		if item.QuantityRequested <= 0 {
			return true
		}
	}
	return false
}

func (o *Order) receivedAll() bool {
	for _, item := range o.List {
		if item.Status != Filled {
			return false
		}
	}
	return true
}

func (o *Order) alreadySent() bool {
	return o.Statuses[len(o.Statuses)-1].Type != New
}

func (o *Order) updateStatus(s OrderStatus) {
	o.Statuses = append(o.Statuses, newStatus(s))
}

type PurchaseOrder struct {
	PONumber string
	Supplier string
}

func newPO(number string, supplier string) PurchaseOrder {
	return PurchaseOrder{
		PONumber: number,
		Supplier: supplier,
	}
}

func (o *Order) AddOrderPO(number string, supplier string) {
	_, err := o.findPO(number)
	if err == nil {
		log.Println(ErrPOalreadyExists)
		return
	}
	o.POs = append(o.POs, newPO(number, supplier))
}

func (o *Order) RemoveOrderPO(number string) {
	i, err := o.findPO(number)
	if err != nil {
		log.Println(ErrPOnotFound)
		return
	}
	o.POs = append(o.POs[:i], o.POs[i+1:]...)
}

func (o *Order) UpdateItemPO(id ProductID, number string, supplier string) {
	i, err := o.findItem(id)
	if err != nil {
		log.Println(ErrPOnotFound)
		return
	}
	o.List[i].updatePO(number, supplier)
}

func (o *Order) RemoveItemPO(id ProductID) {
	i, err := o.findItem(id)
	if err != nil {
		log.Println(ErrPOnotFound)
	}
	o.List[i].removePO()
}

func (o *Order) findPO(number string) (int, error) {
	for i, po := range o.POs {
		if po.PONumber == number {
			return i, nil
		}
	}
	return 0, ErrItemNotFound
}

type Status struct {
	Date time.Time
	Type OrderStatus
}

func newStatus(s OrderStatus) Status {
	return Status{
		Date: time.Time{},
		Type: s,
	}
}

type OrderStatus int

const (
	New OrderStatus = iota
	Sent
	Received
	Processing
	OnRoute
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case New:
		return "New"
	case Sent:
		return "Sent"
	case Received:
		return "Received"
	case Processing:
		return "Processing"
	case OnRoute:
		return "OnRoute"
	case Complete:
		return "Complete"
	}
	return ""
}
