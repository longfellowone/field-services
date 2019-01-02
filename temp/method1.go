package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// A unique order
type Order struct {
	OrderID  int
	LineItem Items
}

// A slice of order items
type Items struct {
	Items []Item
}

// An individual order item
type Item struct {
	Number string
	Price  int
}

// Adjust the price of a single order item
func (o *Order) AdjustItemPrice(number string, price int) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	for i, item := range o.LineItem.Items {
		if item.Number == number {
			o.LineItem.Items[i].Price = price
			return nil
		}
	}
	return nil
}

func main() {
	order := BuildSampleOrder()
	item := &order.LineItem.Items[0]

	fmt.Printf("Old price: %d\n", item.Price)

	err := order.AdjustItemPrice("Num1", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("New price: %d\n", item.Price)
}

// Build sample order
func BuildSampleOrder() *Order {
	return &Order{
		OrderID: 1,
		LineItem: Items{
			Items: []Item{
				{
					Number: "Num1",
					Price:  50,
				},
				{
					Number: "Num2",
					Price:  40,
				},
			},
		},
	}
}
