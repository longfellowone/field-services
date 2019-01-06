package main

import (
	"context"
	"field/pkg/ordering"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

const (
	inMemory = false
)

func main() {
	var service *ordering.Service

	if inMemory {
		service = initializeFieldServicesInMemory()
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, "mongodb://default:password@localhost:27017")
		if err != nil {
			log.Fatal(err)
		}
		defer cancel()
		//err = client.Disconnect(ctx)

		db := client.Database("field")
		service = initializeFieldServices(db)
	}

	service.CreateOrder("oid1", "pid1")

	//service.CreateOrder("oid2", "pid1")

	order, _ := service.FindOrder("oid1")
	fmt.Println(order)
	//
	//order.AddItem("uuid1", "name1")
	//order.AddItem("uuid2", "name2")
	//order.AddItem("uuid3", "name3")
	//
	//order.RemoveItem("uuid1")
	//
	//order.UpdateQuantityRequested("uuid2", 40)
	//order.UpdateQuantityRequested("uuid3", 30)
	//order.Send()
	//
	//order.ReceiveItem("uuid2", 40)
	//order.ReceiveItem("uuid3", 30)
	//
	//order.UpdatePO("uuid3", "po3")
	//
	//fmt.Println(order)
}

//
//
//order1.AddItem("3", "name3", orders.FT)
//order1.AddItem("4", "name4", orders.FT)
//order1.RemoveItem("4")
//
//order2.AddItem("1", "name1", orders.FT)
//order2.AddItem("2", "name2", orders.FT)
//
//order2.UpdateQuantityRequested("1", 12)
//order2.UpdateQuantityRequested("2", 23)
//
//order2.SendOrder()
//
//order2.AddOrderPO("po1", "s1")
//order2.AddOrderPO("po2", "s2")
//order2.AddOrderPO("po3", "s3")
//
//order2.RemoveOrderPO("po2")
//
//order2.UpdateItemPO("1", "po4", "s4")
//order2.UpdateItemPO("1", "po5", "s5")
//order2.UpdateItemPO("2", "po7", "s7")
//order2.UpdateItemPO("2", "po6", "s6")
//
//order2.RemoveItemPO("1")
//
//order2.UpdateItemPO("1", "po9", "s9")
//
//order2.RemoveItemPO("2")
//
//order2.ReceiveQuantity("1", 12)
//order2.ReceiveQuantity("2", 23)
//
//order1.UpdateQuantityRequested("3", 84)
//order1.ReceiveQuantity("3", 73)
//
//for _, v := range orders {
//fmt.Printf("[OID]: %v - [PID]: %v - [STATUS]: ", v.OrderUUID, v.ProjectUUID)
//for _, v := range v.Statuses {
//fmt.Printf("->%v ", v.Type)
//}
//fmt.Printf("[POs]: ")
//for _, po := range v.POs {
//fmt.Printf("%v | ", po.PONumber)
//}
//fmt.Println("")
//for _, v := range v.List {
//fmt.Printf("\t%v %v %v(%v) req:%v rec:%v po:%v\n", v.ProductID, v.Name, v.Status, v.LastUpdate.Format("3:04PM"), v.QuantityRequested, v.QuantityReceived, v.PO.PONumber)
////Mon Jan 2 15:04:05 MST 2006
//}
//}
