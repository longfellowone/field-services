package test

import "fmt"

type Order struct {
	OrderID  int
	LineItem []Item
}

type Item struct {
	Number string
	Price  int
}

func (o *Order) AdjustItemPrice(number string, price int) {
	for i, v := range o.LineItem {
		if v.Number == number {
			o.LineItem[i].Price = price
		}
	}
}

func main() {
	order := Order{
		OrderID: 1,
		LineItem: []Item{{
			Number: "Num1",
			Price:  50,
		}, {
			Number: "Num2",
			Price:  40,
		}},
	}

	fmt.Println(order.LineItem[0])
	order.AdjustItemPrice("Num1", 0)
	fmt.Println(order.LineItem[0])
}
