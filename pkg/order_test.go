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
		OrderUUID:   "cc80e4ba-79ec-42e5-8f85-d46bff29a7d6",
		ProjectUUID: "259d9ebc-1080-40e0-8e2d-3bbeec82dcb8",
	}

	if have.OrderUUID != want.OrderUUID {
		t.Errorf("have %v\n want %v\n", have, want)
	}
}

type AddItem struct {
	id   supply.ProductUUID
	name string
	uom  supply.UOM
}

var addItem = []AddItem{
	{supply.ProductUUID("cc80e4ba-79ec-42e5-8f85-d46bff29a7d6"), "Marker", supply.EA},
	{supply.ProductUUID("a19ca654-db0b-450f-8b89-2a24a910bf7d"), "Pencil", supply.EA},
	{supply.ProductUUID("e07b043e-e4cd-40d0-aefc-32a52ee7edda"), "1/2\" EMT Conduit", supply.FT},
}

func TestOrder_AddItem(t *testing.T) {
	order := supply.Create("", "")

	for _, test := range addItem {
		order.AddItem(test.id, test.name, test.uom)
	}

	if len(order.MaterialList.Items) != len(addItem) {
		t.Errorf("AddItem(): Items not added to list\nHave: %v", order)
	}
}

//func TestOrder_AddItem(t *testing.T) {
//	order := orderWithOneItem
//	order.AddItem(supply.ProductUUID("1ed57fbc-230b-4730-9766-a26235efe79b"), "Pencil2", supply.UOM(supply.EA))
//	want := orderWithTwoItems
//
//	if !reflect.DeepEqual(want, order) {
//		t.Errorf("\nHave %v\nWant %v", order, want)
//	}
//}

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

func TestOrder_ReceiveItem(t *testing.T) {

	testCases := []struct {
		name        string
		shouldError bool
		requested   uint
		received    uint
		remaining   uint
		itemStatus  supply.ItemStatus
		orderStatus supply.OrderStatus
	}{
		{
			name:        "receive all items requested",
			shouldError: false,
			requested:   40,
			received:    39,
			remaining:   0,
		},
	}

	item := supply.Item{
		ProductUUID: "c3ba26e7-ed1f-4119-b7dd-cf1a8de6dc46",
	}

	order := supply.Order{
		OrderHistory: []supply.Event{{
			OrderStatus: supply.Sent,
		}},
	}

	//testOrderStatus := func(t *testing.T, have supply.OrderStatus) {
	//	if have != supply.Complete {
	//		t.Errorf("want: [%v] have: [%v]", supply.Complete, have)
	//	}
	//}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			order.ReceiveItem(item.ProductUUID, test.received)

			if test.received >= test.requested {
			}
		})
	}

	//if len(order.MaterialList.Items) != len(addItem) {
	//	t.Errorf("AddItem(): Items not added to list\nHave: %v", order)
	//}
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
