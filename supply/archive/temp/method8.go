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

	fmt.Printf("Old price: %d\n", item.Quantity)

	err := order.AdjustItemQuantity("Num1", 1)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("New price: %d\n", item.Quantity)
}

// A unique order
type Order struct {
	OrderID int
	MaterialList
}

//Adjust the price of a single order item
func (o *Order) AdjustItemQuantity(id ProductUUID, quantity int) (err error) {
	o.MaterialList, err = o.adjustItemQuantity(id, quantity)
	return
}

// A slice of order items
type MaterialList struct {
	Items []Item
}

// Range threw items and adjust price of item number
func (m MaterialList) adjustItemQuantity(id ProductUUID, quantity int) (MaterialList, error) {
	if quantity <= 0 {
		return m, ErrQuantityZero
	}

	i, err := m.findItem(id)
	if err != nil {
		return m, err
	}
	m.Items[i] = m.Items[i].adjustQuantity(quantity)

	return m, nil
}

func (m MaterialList) findItem(id ProductUUID) (int, error) {
	for i, item := range m.Items {
		if item.ProductUUID == id {
			return i, nil
		}
	}
	return 0, ErrItemNotFound
}

type ProductUUID string

// An individual item
type Item struct {
	ProductUUID ProductUUID
	Quantity    int
}

// Adjust price of an item
func (i Item) adjustQuantity(quantity int) Item {
	i.Quantity = quantity
	return i
}

// Build sample order
func BuildSampleOrder() *Order {
	items := []Item{
		{
			ProductUUID: "Num1",
			Quantity:    50,
		},
		{
			ProductUUID: "Num2",
			Quantity:    60,
		}}

	return &Order{
		OrderID: 0,
		MaterialList: MaterialList{
			Items: items,
		},
	}
}
