package main

import "fmt"

func main() {
	order := BuildSampleOrder()
	item := &order.LineItem[0]

	fmt.Println("start:", item)

	order.AdjustItemPricev1("Num1", 0)
	fmt.Println("   v1:", item)

	order.AdjustItemPricev2("Num1", 10)
	fmt.Println("   v2:", item)
}

type Order struct {
	OrderID  int
	LineItem []Item
}

// V1: Modifies the value directly
func (o *Order) AdjustItemPricev1(number string, price int) {
	for i, v := range o.LineItem {
		if v.Number == number {
			o.LineItem[i].Price = price
		}
	}
}

// V2: Modifies the value using the Item method adjustPrice
func (o *Order) AdjustItemPricev2(number string, price int) {
	for i, v := range o.LineItem {
		if v.Number == number {
			o.LineItem[i].adjustPrice(price)
		}
	}
}

type Item struct {
	Number string
	Price  int
}

func (i *Item) adjustPrice(price int) {
	i.Price = price
}

func BuildSampleOrder() *Order {
	return &Order{
		OrderID: 1,
		LineItem: []Item{{
			Number: "Num1",
			Price:  50,
		}, {
			Number: "Num2",
			Price:  40,
		}},
	}
}
