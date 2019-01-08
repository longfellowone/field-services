// https://play.golang.org/p/8npy-gWb45e
package main

import (
	"fmt"
)

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

// A unique order
type Order struct {
	OrderID  int
	LineItem Items
}

// Adjust the price of a single order item
func (o *Order) AdjustItemPrice(number string, price int) error {
	if price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	for i := range o.LineItem.Items {
		item := &o.LineItem.Items[i]

		if item.Number == number {
			item.adjustPrice(price)
			return nil
		}
	}
	return nil
}

// A slice of order items
type Items struct {
	Items []Item
}

// An individual item
type Item struct {
	Number string
	Price  int
}

// Adjust price of an item
func (i *Item) adjustPrice(price int) {
	i.Price = price
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
