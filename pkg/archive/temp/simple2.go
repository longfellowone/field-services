package main

import (
	"errors"
)

type Order struct {
	OrderID int
	Items   []Item
}

func (o *Order) AdjustItemPrice(number string, price int) error {
	if price <= 0 {
		return errors.New("price cannot be 0")
	}
	i, err := o.findItem(number)
	if err != nil {
		return err
	}
	o.Items[i].adjustPrice(price)
	return nil
}

func (o *Order) RemoveItem(number string) error {
	i, err := o.findItem(number)
	if err != nil {
		return err
	}
	o.removeItem(i)
	return nil
}

func (o *Order) AddItem(number string, price int) {
	o.Items = append(o.Items, newItem(number, price))
}

func (o *Order) removeItem(i int) {
	o.Items = append(o.Items[:i], o.Items[i+1:]...)
}

func (o *Order) findItem(number string) (int, error) {
	for i, item := range o.Items {
		if item.Number == number {
			return i, nil
		}
	}
	return 0, errors.New("item not found")
}

type Item struct {
	Number string
	Price  int
}

func (i *Item) adjustPrice(price int) {
	i.Price = price
}

func newItem(number string, price int) Item {
	return Item{
		Number: number,
		Price:  price,
	}
}
