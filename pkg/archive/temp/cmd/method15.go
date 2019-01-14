package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	order := BuildSampleList()

	fmt.Println(order)

	fmt.Printf("Old price: %d\n", order.Items[0].quantityRequested)

	order.AdjustQuantityRequested("num1", 1)
	order.AddItem("num3")

	fmt.Printf("New price: %d\n", order.Items[0].quantityRequested)

	fmt.Println(order)
}

type Order struct {
	orderID   string
	projectID string
	materialList
}

func Create(orderid string, projectid string) *Order {
	return &Order{
		orderID:   orderid,
		projectID: projectid,
		materialList: materialList{
			Items: nil,
		},
	}
}

type materialList struct {
	Items []Item
}

func (m *materialList) AdjustQuantityRequested(id string, qr int) {
	item, err := m.findItem(id)
	if err != nil {
		log.Fatal(err)
	}
	item.updateRequested(qr)
}

func (m *materialList) AddItem(id string) {
	m.Items = append(m.Items, newItem(id))
}

func (m *materialList) findItem(id string) (*Item, error) {
	for i := range m.Items {
		if m.Items[i].productUUID == id {
			return &m.Items[i], nil
		}
	}
	return &Item{}, errors.New("item not found")
}

// An individual item
type Item struct {
	productUUID       string
	quantityRequested int
}

func newItem(id string) Item {
	return Item{
		productUUID: id,
	}
}

func (i *Item) updateRequested(qr int) {
	i.quantityRequested = qr
}

// Build sample order
func BuildSampleList() *Order {
	items := []Item{
		{
			productUUID:       "num1",
			quantityRequested: 50,
		},
		{
			productUUID:       "num2",
			quantityRequested: 100,
		},
	}

	return &Order{
		orderID:   "order1",
		projectID: "project1",
		materialList: materialList{
			Items: items,
		},
	}
}
