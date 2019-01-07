package temp

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrQuantityZero = errors.New("price must be greater than 0")
	ErrItemNotFound = errors.New("item not found")
)

func main() {
	order := BuildSampleOrder()

	fmt.Println(order.MaterialList)
	order.AddItem("Num3")

	fmt.Printf("Old price: %d\n", order.MaterialList[0].QuantityRequested)

	err := order.AdjustItemQuantity("Num1", 1)
	if err != nil {
		log.Println(err)
	}

	//order.Send()

	fmt.Printf("New price: %d\n", order.MaterialList[0].QuantityRequested)
	fmt.Println(order.MaterialList)
}

// A unique order
type Order struct {
	OrderID int
	MaterialList
	OrderStatus
}

//func (o *Order) AddItem(id productUUID) {
//	item := newItem(id)
//	o.materialLists = append(o.materialLists, item)
//}

//func (o *Order) RemoveItem(id productUUID) error {
//	item := o.findItem(id)
//	if item < 0 {
//		return ErrItemNotFound
//	}
//	o.materialLists = append(o.materialLists.Items[:item], o.materialLists[item+1:]...)
//	return nil
//}

func (o *Order) Send() {
	//status := itemStatus{
	//	Date: "1",
	//	Type: "",
	//}
	//
	//fmt.Println(status)
}

func (o *Order) AddItem(id ProductUUID) {
	o.MaterialList = o.MaterialList.addItem(id)
}

//func (o *Order) AdjustItemQuantity(id productUUID, qr quantityRequested) (err error) {
//	err = o.materialLists.adjustItemQuantity(id, qr)
//	if err != nil {
//		log.Println(err)
//	}
//	return nil
//}

//Adjust the price of a single order item
//func (o *Order) AdjustItemQuantity(id productUUID, quantity int) (err error) {
//	o.materialLists, err = o.adjustItemQuantity(id, quantity)
//	return
//}

// A slice of order items
type MaterialList []Item

// Range threw items and adjust price of item number
func (m MaterialList) AdjustItemQuantity(id ProductUUID, qr QuantityRequested) error {
	if qr <= 0 {
		return ErrQuantityZero
	}

	item := m.findItem(id)
	if item < 0 {
		return ErrItemNotFound
	}

	m[item].QuantityRequested = qr
	//m[item] = m[item].adjustQuantity(qr)

	return nil
}

func (m MaterialList) addItem(id ProductUUID) MaterialList {
	return append(m, newItem(id))
}

func (m MaterialList) findItem(id ProductUUID) int {
	for item := range m {
		if m[item].ProductUUID == id {
			return item
		}
	}
	return -1
}

type OrderStatus struct {
	Statuses []Status
}

type Status struct {
	Date string
	Type string
}

type ProductUUID string
type QuantityRequested int
type QuantityReceived int

// An individual item
type Item struct {
	ProductUUID
	QuantityRequested
	QuantityReceived
	Status
}

func newItem(id ProductUUID) Item {
	return Item{
		ProductUUID:       id,
		QuantityRequested: 0,
		QuantityReceived:  0,
	}
}

// Adjust price of an item
func (i Item) adjustQuantity(qr QuantityRequested) Item {
	i.QuantityRequested = qr
	return i
}

// Build sample order
func BuildSampleOrder() *Order {
	items := []Item{
		{
			ProductUUID:       "Num1",
			QuantityRequested: 50,
			QuantityReceived:  0,
		},
		{
			ProductUUID:       "Num2",
			QuantityRequested: 100,
			QuantityReceived:  0,
		},
	}

	return &Order{
		OrderID:      0,
		MaterialList: items,
		OrderStatus: OrderStatus{
			Statuses: nil,
		},
	}
}
