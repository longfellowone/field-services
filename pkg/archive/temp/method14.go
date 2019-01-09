package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

var (
	ErrQuantityZero = errors.New("price must be greater than 0")
	ErrItemNotFound = errors.New("item not found")
)

func main() {
	order := BuildSampleOrder()

	fmt.Println(order.materialLists)

	fmt.Printf("Old price: %d\n", order.materialLists[0].quantityRequested)

	err := AddItem("Num3", order)
	if err != nil {
		log.Fatal(err)
	}

	err = AddItem("Num4", order)
	if err != nil {
		log.Fatal(err)
	}
	//order.ModifyQuantityRequested("Num2", 3)
	//order.ModifyQuantityRequested("Num3", 7)
	//order.RemoveItem("Num1")
	//
	//order.ReceiveItem("Num2", 1)

	//order.Send()

	fmt.Printf("New price: %d\n", order.materialLists[0].quantityRequested)
	fmt.Println(order.materialLists)
}

type Order struct {
	orderUUID     string
	projectUUID   string
	materialLists []Item
	dateOrdered   time.Time
	status        string
}

func Create(id string, pid string) Order {
	return Order{
		orderUUID:   id,
		projectUUID: pid,
	}
}

func (o *Order) Send() {

}

func AddItem(id string, o *Order) error {
	_, err := o.findItem(id)
	if err == nil {
		return ErrItemAlreadyExists
	}
	o.materialLists = append(o.materialLists, newItem(id))
	return nil
	// Test for len() then test [len()-1}
}

func (o *Order) RemoveItem(id string) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.materialLists = append(o.materialLists[:i], o.materialLists[i+1:]...)
	return nil
}

func (o *Order) ModifyQuantityRequested(id string, quantity int) error {
	if quantity <= 0 {
		return ErrQuantityZero
	}
	err := o.updateItem(id, modifyQuantityRequested(quantity))
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) ReceiveItem(id string, quantity int) error {
	if quantity <= 0 {
		return ErrQuantityZero
	}
	err := o.updateItem(id, receiveItem(quantity))
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) findItem(id string) (int, error) {
	for i := range o.materialLists {
		if o.materialLists[i].productUUID == id {
			return i, nil
		}
	}
	return 0, ErrItemNotFound
}

type ItemOption func(i *Item)

func (o *Order) updateItem(id string, opt ItemOption) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	opt(&o.materialLists[i])
	return nil
}

func modifyQuantityRequested(quantity int) ItemOption {
	return func(i *Item) {
		i.quantityRequested = quantity
	}
}

func receiveItem(quantity int) ItemOption {
	return func(i *Item) {
		i.quantityReceived = quantity
		switch {
		case i.quantityReceived == i.quantityRequested:
			i.quantityRemaining = 0
			//i.itemStatus = Filled
		case quantity > 0 && quantity < i.quantityRequested:
			i.quantityRemaining = i.quantityRequested - i.quantityReceived
			//i.itemStatus = BackOrdered
		case i.quantityReceived > i.quantityRequested:
			i.quantityRemaining = 0
			//i.itemStatus = OrderExceeded
		}
	}
}

// An individual item
type Item struct {
	productUUID       string
	quantityRequested int
	quantityReceived  int
	quantityRemaining int
	itemStatus        string
}

func newItem(id string) Item {
	return Item{
		productUUID:       id,
		quantityRequested: 0,
		quantityReceived:  0,
	}
}

// Build sample order
func BuildSampleOrder() *Order {
	items := []Item{
		{
			productUUID:       "Num1",
			quantityRequested: 50,
			quantityReceived:  0,
		},
		{
			productUUID:       "Num2",
			quantityRequested: 100,
			quantityReceived:  0,
		},
	}

	return &Order{
		orderUUID:     "orderid1",
		materialLists: items,
	}
}
