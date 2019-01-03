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

	var fs *ordering.Service

	if inMemory {
		fs = initializeFieldServicesInMemory()
	}
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

	fs.CreateNewOrder("oid1", "pid1")
	fs.CreateNewOrder("oid2", "pid1")

	result3, _ := fs.FindAllProjectOrders("pid1")

	_ = result3[0].AddItem("3", "name3", material.FT)
	_ = result3[0].AddItem("4", "name4", material.FT)
	_ = result3[0].AddItem("3", "name3", material.FT)
	err := result3[0].RemoveItem("4")
	if err != nil {
		fmt.Println(err)
	}
	//err = result3[0].RemoveItemFromList("4")
	//if err != nil {
	//	fmt.Println(err)
	//}
	_ = result3[1].AddItem("1", "name1", material.FT)
	_ = result3[1].AddItem("1", "name1", material.FT)
	_ = result3[1].AddItem("2", "name2", material.FT)
	//
	_ = result3[1].UpdateQuantityRequested("1", 12)
	_ = result3[1].UpdateQuantityRequested("2", 23)
	//
	_ = result3[1].SendOrder()

	_ = result3[1].AddOrderPO("po1", "s1")
	_ = result3[1].AddOrderPO("po2", "s2")
	_ = result3[1].AddOrderPO("po3", "s3")
	//
	_ = result3[1].RemoveOrderPO("po2")
	//
	_ = result3[1].UpdateItemPO("1", "po4", "s4")
	_ = result3[1].UpdateItemPO("1", "po5", "s5")
	_ = result3[1].UpdateItemPO("2", "po7", "s7")
	_ = result3[1].UpdateItemPO("2", "po6", "s6")
	//
	_ = result3[1].RemoveItemPO("1")
	//
	_ = result3[1].UpdateItemPO("1", "po9", "s9")
	//
	_ = result3[1].RemoveItemPO("2")

	_ = result3[1].ReceiveQuantity("1", 12)
	_ = result3[1].ReceiveQuantity("2", 23)

	_ = result3[0].UpdateQuantityRequested("3", 84)
	_ = result3[0].ReceiveQuantity("3", 73)

	//fmt.Println("Find by project")
	for _, v := range result3 {
		fmt.Printf("[OID]: %v - [PID]: %v - [STATUS]: ", v.OrderID, v.ProjectID) //[len(v.Statuses)-1].Type
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
