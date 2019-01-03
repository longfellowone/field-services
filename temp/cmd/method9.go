package main

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
	item := &order.MaterialList.Items[0]

	order.AddItem("Num3")
	fmt.Println(order.MaterialList.Items)

	fmt.Printf("Old price: %d\n", item.QuantityRequested)

	err := order.AdjustItemQuantity("Num1", 1)
	if err != nil {
		log.Println(err)
	}

	//order.Send()

	fmt.Printf("New price: %d\n", item.QuantityRequested)
}

// A unique order
type Order struct {
	OrderID      int
	MaterialList MaterialList
	OrderStatus
}

//func (o *Order) AddItem(id ProductUUID) {
//	item := newItem(id)
//	o.MaterialList = append(o.MaterialList, item)
//}

//func (o *Order) RemoveItem(id ProductUUID) error {
//	item := o.findItem(id)
//	if item < 0 {
//		return ErrItemNotFound
//	}
//	o.MaterialList = append(o.MaterialList.Items[:item], o.MaterialList[item+1:]...)
//	return nil
//}

func (o *Order) Send() {
	//status := Status{
	//	Date: "1",
	//	Type: "",
	//}
	//
	//fmt.Println(status)
}

func (o *Order) AddItem(id ProductUUID) {
	o.MaterialList = o.MaterialList.addItem(id)
}

func (o *Order) AdjustItemQuantity(id ProductUUID, qr QuantityRequested) (err error) {
	list, err := o.MaterialList.adjustItemQuantity(id, qr)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(list)
	return nil
}

//Adjust the price of a single order item
//func (o *Order) AdjustItemQuantity(id ProductUUID, quantity int) (err error) {
//	o.MaterialList, err = o.adjustItemQuantity(id, quantity)
//	return
//}

// A slice of order items
type MaterialList struct {
	Items []Item
}

func (m MaterialList) addItem(id ProductUUID) MaterialList {
	item := newItem(id)
	m.Items = append(m.Items, item)
	return m
}

// Range threw items and adjust price of item number
func (m MaterialList) adjustItemQuantity(id ProductUUID, qr QuantityRequested) (MaterialList, error) {
	if qr <= 0 {
		return MaterialList{}, ErrQuantityZero
	}

	item := m.findItem(id)
	if item < 0 {
		return MaterialList{}, ErrItemNotFound
	}

	m.Items[item] = m.Items[item].adjustQuantity(qr)
	fmt.Println(true)

	return m, nil
}

func (m MaterialList) findItem(id ProductUUID) int {
	for item := range m.Items {
		if m.Items[item].ProductUUID == id {
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
		OrderID: 0,
		MaterialList: MaterialList{
			Items: items,
		},
		OrderStatus: OrderStatus{
			Statuses: nil,
		},
	}
}
