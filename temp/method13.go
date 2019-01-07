package main

import (
	"errors"
	"fmt"
	"time"
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
	order.ModifyQuantityRequested("Num2", 3)
	order.ModifyQuantityRequested("Num3", 7)
	order.RemoveItem("Num1")

	order.ReceiveItem("Num2", 1)

	//order.Send()

	fmt.Printf("New price: %d\n", order.MaterialList[0].QuantityRequested)
	fmt.Println(order.MaterialList)
}

type Order struct {
	OrderID string
	ProjectID string
	MaterialList []Item
	LastChange time.Time
	DateOrdered time.Time
}

func Create(id string, pid string) *Order {
	return &Order{
		OrderID:      id,
		ProjectID:    pid,
		MaterialList: nil,
		LastChange:   time.Time{},
		DateOrdered:  time.Time{},
	}
}

func (o *Order) Send() {

}

func (o *Order) AddItem(id string) error {
	_, err := o.findItem(id)
	if err == nil {
		return ErrItemAlreadyExists
	}
	o.MaterialList = append(o.MaterialList, newItem(id))
	return nil
}

func (o *Order) RemoveItem(id string) error {
	i, err := o.findItem(id)
	if err != nil {
		return err
	}
	o.MaterialList = append(o.MaterialList[:i], o.MaterialList[i+1:]...)
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
	for i := range o.MaterialList {
		if o.MaterialList[i].ProductUUID == id {
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
	opt(&o.MaterialList[i])
	return nil
}

func modifyQuantityRequested(quantity int) ItemOption {
	return func(i *Item) {
		i.QuantityRequested = quantity
	}
}

func receiveItem(quantity int) ItemOption {
	return func(i *Item) {
		i.QuantityReceived = quantity
		switch {
		case i.QuantityReceived == i.QuantityRequested:
			//i.ItemStatus = Filled
			i.QuantityRemaining = 0
		case quantity > 0 && quantity < i.QuantityRequested:
			//i.ItemStatus = BackOrdered
			i.QuantityRemaining = i.QuantityRequested - i.QuantityReceived
		case i.QuantityReceived > i.QuantityRequested:
			//i.ItemStatus = OrderExceeded
			i.QuantityRemaining = 0
		}
	}
}
// An individual item
type Item struct {
	ProductUUID string
	QuantityRequested int
	QuantityReceived int
	QuantityRemaining int
	ItemStatus string
}

func newItem(id string) Item {
	return Item{
		ProductUUID:       id,
		QuantityRequested: 0,
		QuantityReceived:  0,
	}
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
		OrderID:      "orderid1",
		MaterialList: items,
	}
}
