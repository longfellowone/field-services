// https://play.golang.org/p/nlpUz9Acqxl
package test

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

	for i, item := range o.LineItem.Items {
		if item.Number == number {
			o.LineItem.Items[i].Price = price
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
