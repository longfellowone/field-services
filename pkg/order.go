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
	ErrPOalreadyExists   = errors.New("PO already exists")
	ErrPOnotFound        = errors.New("PO not found")
	ErrItemAlreadyExists = errors.New("item already exists")
	ErrOrderComplete     = errors.New("order complete cannot add item")
)

type OrderRepository interface {
	Save(o *Order) error
	Find(id OrderUUID) (*Order, error)
	FindAllFromProject(id ProjectUUID) ([]*Order, error)
}

type OrderUUID string
type ProjectUUID string

type Order struct {
	OrderUUID
	ProjectUUID
	MaterialList
	PurchaseOrders
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
	if !o.okToSend() {
		return
	}
	o.newEvent(Sent)
}

func (o *Order) AddItem(uuid ProductUUID, name string) {
	if o.lastEvent() != Created {
		log.Println(ErrOrderSent)
		return
	}
	o.MaterialList = o.addItem(uuid, name)
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
}

func (o *Order) okToSend() bool {
	for i := range o.MaterialList.Items {
		switch {
		case o.MaterialList.Items[i].QuantityRequested == 0:
			log.Println(ErrQuantityZero)
			return false
		}
	}
	return true
}

func (o *Order) lastEvent() OrderStatus {
	return o.OrderHistory[len(o.OrderHistory)-1].OrderStatus
}

func (o *Order) newEvent(event OrderStatus) {
	o.OrderHistory = append(o.OrderHistory, createEvent(event))
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

//func (o *Order) UpdateMaterialList(item Item) error {
//	switch {
//	case o.lastEvent() == Created:
//		o.updateMaterialList(item)
//	case o.lastEvent() == Sent:
//		return ErrOrderSent
//	case o.lastEvent() == OnRoute:
//		fmt.Println(item)
//		//o.ReceiveItem(item)
//	case o.lastEvent() == Complete:
//		return ErrOrderComplete
//	}
//	return nil
//}

//func (o *Order) updateMaterialList(item Item) {
//	switch {
//	case item.QuantityRequested < 0:
//		o.MaterialList = o.MaterialList.removeItem(item.ProductUUID)
//	//case item:
//	//	o.MaterialList
//	case item.QuantityReceived == 0 || item.ProductUUID != "":
//		o.MaterialList = append(o.MaterialList, item)
//	}
//}
