package supply

import (
	"errors"
	"log"
	"time"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrOrderSent         = errors.New("order already sent")
	ErrMustHaveItems     = errors.New("order must have at least 1 item")
	ErrQuantityZero      = errors.New("quantity of all order items must be greater than 0")
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyOnList = errors.New("item already on list")
	ErrItemQuantityZero  = errors.New("item quantity must be greater than 0")
)

type OrderRepository interface {
	Save(o *Order) error
	Find(uuid OrderUUID) (*Order, error)
	FindAllFromProject(uuid ProjectUUID) ([]Order, error)
}

type OrderUUID string
type ProjectUUID string

type Order struct {
	OrderUUID
	ProjectUUID
	MaterialList
	OrderHistory []Event
}

func Create(id OrderUUID, pid ProjectUUID) *Order {
	event := createEvent(Created)

	return &Order{
		OrderUUID:    id,
		ProjectUUID:  pid,
		OrderHistory: []Event{event},
	}
}

func (o *Order) Send() {
	switch {
	case o.MaterialList.Items == nil:
		log.Println(ErrMustHaveItems)
		return
	case o.missingQuantities():
		log.Println(ErrQuantityZero)
		return
	case o.alreadySent():
		log.Println(ErrOrderSent)
		return
	}
	o.newEvent(Sent)
}

func (o *Order) AddItem(uuid ProductUUID, name string, uom UOM) {
	_, item := o.findItem(uuid)

	if item.ProductUUID != "" {
		log.Println(ErrItemAlreadyOnList)
		return
	}
	if o.lastEvent() != Created {
		log.Println(ErrOrderSent)
		return
	}

	o.MaterialList.Items = append(o.MaterialList.Items, newItem(uuid, name, uom))
}

func (o *Order) updateList() {

}

func (o *Order) RemoveItem(uuid ProductUUID) {
	if o.lastEvent() != Created {
		log.Println(ErrOrderSent)
		return
	}
	o.MaterialList = o.removeItem(uuid)
}

func (o *Order) ReceiveItem(uuid ProductUUID, quantity uint) {
	o.receiveItem(uuid, quantity)
	if o.receivedAll() {
		o.newEvent(Complete)
	}
}

func (o *Order) missingQuantities() bool {
	for i := range o.MaterialList.Items {
		if o.MaterialList.Items[i].QuantityRequested == 0 {
			return true
		}
	}
	return false
}

func (o *Order) newEvent(event OrderStatus) {
	o.OrderHistory = append(o.OrderHistory, createEvent(event))
}

func (o *Order) lastEvent() OrderStatus {
	return o.OrderHistory[len(o.OrderHistory)-1].OrderStatus
}

func (o *Order) alreadySent() bool {
	return o.OrderHistory[len(o.OrderHistory)-1].OrderStatus != Created
}

type Event struct {
	Date time.Time
	OrderStatus
}

func createEvent(status OrderStatus) Event {
	return Event{
		Date:        time.Now(),
		OrderStatus: status,
	}
}

type OrderStatus int

const (
	Created OrderStatus = iota
	Sent
	OnRoute
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case Created:
		return "New"
	case Sent:
		return "Sent"
	case OnRoute:
		return "OnRoute"
	case Complete:
		return "Complete"
	}
	return ""
}
