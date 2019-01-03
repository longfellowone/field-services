package main

import (
	"field/pkg"
	"field/pkg/ordering"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	//defaultGRPCPort = 9090
	//defaultDBHost   = "localhost"
	//defaultDBPort   = 5432
	//defaultDBName   = "default"
	//defaultDBUser   = "default"
	//defaultDBPasswd = "password"
	//sslMode         = "disable"
	inMemory = true
)

func main() {

	// flags := flag.Parse()

	//var (
	//	dbHost                   = defaultDBHost
	//	dbPort                   = defaultDBPort
	//	dbUser                   = defaultDBUser
	//	dbPasswd                 = defaultDBPasswd
	//	dbName                   = defaultDBName
	//	postgresConnectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPasswd, dbName, sslMode)
	//)

	var service *ordering.Service

	if inMemory {
		service = initializeFieldServicesInMemory()
	}

	fmt.Println(service)
	//else {
	//	db, err := sql.Open("postgres", postgresConnectionString)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer db.Close()
	//
	//	if err = db.Ping(); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fs = initializeFieldServices(db)
	//}

	service.CreateOrder("oid1", "pid1")
	service.CreateOrder("oid2", "pid1")

	orders, _ := service.FindAllProjectOrders("pid1")
	order1 := orders[0]
	order2 := orders[1]

	order1.AddItem("3", "name3", material.FT)
	order1.AddItem("4", "name4", material.FT)
	order1.RemoveItem("4")

	order2.AddItem("1", "name1", material.FT)
	order2.AddItem("2", "name2", material.FT)

	order2.UpdateQuantityRequested("1", 12)
	order2.UpdateQuantityRequested("2", 23)

	order2.SendOrder()

	order2.AddOrderPO("po1", "s1")
	order2.AddOrderPO("po2", "s2")
	order2.AddOrderPO("po3", "s3")

	order2.RemoveOrderPO("po2")

	order2.UpdateItemPO("1", "po4", "s4")
	order2.UpdateItemPO("1", "po5", "s5")
	order2.UpdateItemPO("2", "po7", "s7")
	order2.UpdateItemPO("2", "po6", "s6")

	order2.RemoveItemPO("1")

	order2.UpdateItemPO("1", "po9", "s9")

	order2.RemoveItemPO("2")

	order2.ReceiveQuantity("1", 12)
	order2.ReceiveQuantity("2", 23)

	order1.UpdateQuantityRequested("3", 84)
	order1.ReceiveQuantity("3", 73)

	for _, v := range orders {
		fmt.Printf("[OID]: %v - [PID]: %v - [STATUS]: ", v.OrderID, v.ProjectID)
		for _, v := range v.Statuses {
			fmt.Printf("->%v ", v.Type)
		}
		fmt.Printf("[POs]: ")
		for _, po := range v.POs {
			fmt.Printf("%v | ", po.PONumber)
		}
		fmt.Println("")
		for _, v := range v.List {
			fmt.Printf("\t%v %v %v(%v) req:%v rec:%v po:%v\n", v.ProductID, v.Name, v.Status, v.LastUpdate.Format("3:04PM"), v.QuantityRequested, v.QuantityReceived, v.PO.PONumber)
			//Mon Jan 2 15:04:05 MST 2006
		}
	}
}
