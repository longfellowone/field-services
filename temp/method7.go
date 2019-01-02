// https://play.golang.org/p/whwQ7lCHcMc
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
func (o *Order) AdjustItemPrice(number string, price int) (err error) {
	o.LineItem, err = o.LineItem.adjustItemPrice(number, price)
	return
}

// A slice of order items
type Items struct {
	Items []Item
}

// Range threw items and adjust price of item number
func (i Items) adjustItemPrice(number string, price int) (Items, error) {
	if price <= 0 {
		return i, fmt.Errorf("price must be greater than 0")
	}

	for n, item := range i.Items {

		if item.Number == number {
			i.Items[n] = i.Items[n].adjustPrice(price)
			return i, nil
		}
	}
	return i, nil
}

// An individual item
type Item struct {
	Number string
	Price  int
}

// Adjust price of an item
func (i Item) adjustPrice(price int) Item {
	i.Price = price
	return i
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
