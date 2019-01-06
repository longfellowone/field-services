package supply_test

import (
	"reflect"
	"supply/pkg"
	"testing"
	"time"
)

var timeNow = time.Now()

func TestCreate(t *testing.T) {
	have := supply.Create("cc80e4ba-79ec-42e5-8f85-d46bff29a7d6", "259d9ebc-1080-40e0-8e2d-3bbeec82dcb8")
	want := &supply.Order{
		OrderUUID: "cc80e4ba-79ec-42e5-8f85-d46bff29a7d6",
		//ProjectUUID: "259d9ebc-1080-40e0-8e2d-3bbeec82dcb8",
	}

	if have.OrderUUID != want.OrderUUID {
		t.Errorf("have %v\n want %v\n", have, want)
	}
}

func TestOrder_AddItem(t *testing.T) {
	order := orderWithOneItem
	order.AddItem(supply.ProductUUID("1ed57fbc-230b-4730-9766-a26235efe79b"), "Pencil2", supply.UOM(supply.EA))
	want := orderWithTwoItems

	if !reflect.DeepEqual(want, order) {
		t.Errorf("\nHave %v\nWant %v", order, want)
	}
}

func TestOrder_RemoveItem(t *testing.T) {
	order := orderWithTwoItems
	order.RemoveItem(supply.ProductUUID("1ed57fbc-230b-4730-9766-a26235efe79b"))
	want := orderWithOneItem

	if !reflect.DeepEqual(want, order) {
		t.Errorf("\nHave %v\nWant %v", order, want)
	}
}

func TestMaterialList_UpdateQuantityRequested(t *testing.T) {
	order := orderWithTwoItems
	order.UpdateQuantityRequested("1ed57fbc-230b-4730-9766-a26235efe79b", 30)

	want := orderRequestQuantity

	if !reflect.DeepEqual(want, order) {
		t.Errorf("\nHave %v\nWant %v", order, want)
	}
}

var orderWithOneItem = supply.Order{
	OrderUUID:   "cc80e4ba-79ec-42e5-8f85-d46bff29a7d6333",
	ProjectUUID: "259d9ebc-1080-40e0-8e2d-3bbeec82dcb8",
	MaterialList: supply.MaterialList{
		Items: []supply.Item{{
			ProductUUID:       "10484a5a-1b60-4442-ba5c-8e306ec863f89",
			Name:              "Pencil1",
			UOM:               supply.EA,
			QuantityRequested: 40,
			QuantityReceived:  0,
			QuantityRemaining: 0,
			ItemStatus:        0,
			PONumber:          "N/A",
		}},
	},
	OrderHistory: []supply.Event{{
		Date:        timeNow,
		OrderStatus: supply.Created,
	}},
}

var orderWithTwoItems = supply.Order{
	OrderUUID:   "cc80e4ba-79ec-42e5-8f85-d46bff29a7d6333",
	ProjectUUID: "259d9ebc-1080-40e0-8e2d-3bbeec82dcb8",
	MaterialList: supply.MaterialList{
		Items: []supply.Item{{
			ProductUUID:       "10484a5a-1b60-4442-ba5c-8e306ec863f89",
			Name:              "Pencil1",
			UOM:               supply.EA,
			QuantityRequested: 40,
			QuantityReceived:  0,
			QuantityRemaining: 0,
			ItemStatus:        0,
			PONumber:          "N/A",
		}, {
			ProductUUID:       "1ed57fbc-230b-4730-9766-a26235efe79b",
			Name:              "Pencil2",
			UOM:               supply.EA,
			QuantityRequested: 0,
			QuantityReceived:  0,
			QuantityRemaining: 0,
			ItemStatus:        0,
			PONumber:          "N/A",
		}},
	},
	OrderHistory: []supply.Event{{
		Date:        timeNow,
		OrderStatus: supply.Created,
	}},
}

var orderRequestQuantity = supply.Order{
	OrderUUID:   "cc80e4ba-79ec-42e5-8f85-d46bff29a7d6333",
	ProjectUUID: "259d9ebc-1080-40e0-8e2d-3bbeec82dcb8",
	MaterialList: supply.MaterialList{
		Items: []supply.Item{{
			ProductUUID:       "10484a5a-1b60-4442-ba5c-8e306ec863f89",
			Name:              "Pencil1",
			UOM:               supply.EA,
			QuantityRequested: 40,
			QuantityReceived:  0,
			QuantityRemaining: 0,
			ItemStatus:        0,
			PONumber:          "N/A",
		}, {
			ProductUUID:       "1ed57fbc-230b-4730-9766-a26235efe79b",
			Name:              "Pencil2",
			UOM:               supply.EA,
			QuantityRequested: 30,
			QuantityReceived:  0,
			QuantityRemaining: 0,
			ItemStatus:        0,
			PONumber:          "N/A",
		}},
	},
	OrderHistory: []supply.Event{{
		Date:        timeNow,
		OrderStatus: supply.Created,
	}},
}
