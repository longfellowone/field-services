package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrQuantityZero      = errors.New("price must be greater than 0")
	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyExists = errors.New("item already exists")
)

func main() {
	order := BuildSampleOrder()

	fmt.Println(order.MaterialList)

	fmt.Printf("Old price: %d\n", order.MaterialList[0].QuantityRequested)

	order.AddItem("Num3")
	order.AddItem("Num4")
	order.ModifyQuantityRequested("Num1", 7)
	order.ModifyQuantityRequested("Num1", 3)

	//order.RemoveItem("Num2")

	//order.Send()

	fmt.Printf("New price: %d\n", order.MaterialList[0].QuantityRequested)
	fmt.Println(order.MaterialList)
}

type OrderID int
type ProjectID int

type Order struct {
	OrderID
	ProjectID
	MaterialList []Item
	OrderStatus
}

func Create(id OrderID, pid ProjectID) *Order {
	return &Order{
		OrderID:      id,
		ProjectID:    pid,
		MaterialList: nil,
		OrderStatus:  nil,
	}
}

//func (o *Order) AddItem(id ProductUUID) {
//	//if o.findItem(id) > 0 {
//	//	log.Println(ErrItemAlreadyExists)
//	//	return
//	//}
//	o.MaterialList = append(o.MaterialList, newItem(id))
//}

func (o *Order) RemoveItem(id ProductUUID) {
	//i := o.findItem(id)
	//if i < 0 {
	//	log.Println(ErrItemNotFound)
	//	return
	//}
	//o.MaterialList = append(o.MaterialList[:i], o.MaterialList[i+1:]...)
}

func (o *Order) Send() {

}

func (o *Order) AddItem(id ProductUUID) {
	o.MaterialList = append(o.MaterialList, newItem(id))
}

func (o *Order) ModifyQuantityRequested(id ProductUUID, quantity QuantityRequested) {
	if quantity <= 0 {
		log.Println(ErrQuantityZero)
		return
	}
	o.updateItem(id, modifyQuantityRequested(quantity))
}

type UpdateOption func(item *Item)

func (o *Order) updateItem(id ProductUUID, update UpdateOption) {
	for i := range o.MaterialList {
		item := &o.MaterialList[i]
		if  item.ProductUUID == id {
			update(item)
		}
	}
	log.Println(ErrItemNotFound)
}

func modifyQuantityRequested(quantity QuantityRequested) func(item *Item) {
	return func(item *Item) {
		item.QuantityRequested = quantity
	}
}

type OrderStatus []Status

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

func receiveItem(quantity QuantityReceived) func(item Item) Item {
	return func(item Item) Item {
		item.QuantityReceived = quantity
		return item
	}
}

//func (i Item) modifyQuantityRequested(quantity QuantityRequested) Item {
//	i.QuantityRequested = quantity
//	return i
//}

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
		OrderStatus:  nil,
	}
}
