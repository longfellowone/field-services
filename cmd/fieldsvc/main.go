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

	//result2, _ := fs.FindOrder("oid1")
	//fmt.Println("Find order")
	//fmt.Printf("[OID]: %v - [PID]: %v\n", result2.OrderID, result2.ProjectID)

	result3, _ := fs.FindAllProjectOrders("pid1")

	_ = result3[0].AddItemToList("3", "name3", material.FT)
	_ = result3[0].AddItemToList("4", "name4", material.FT)
	_ = result3[0].AddItemToList("3", "name3", material.FT)
	_ = result3[0].RemoveItemFromList("4")

	_ = result3[1].AddItemToList("1", "name1", material.FT)
	_ = result3[1].AddItemToList("1", "name1", material.FT)
	_ = result3[1].AddItemToList("2", "name2", material.FT)

	_ = result3[1].UpdateQuantityRequested("1", 12)
	_ = result3[1].UpdateQuantityRequested("2", 8)

	_ = result3[1].ReceiveQuantity("1", 12)

	_ = result3[1].SendOrder()

	//fmt.Println("Find by project")
	for _, v := range result3 {
		fmt.Printf("[OID]: %v - [PID]: %v - [STATUS]: %v\n", v.OrderID, v.ProjectID, v.Statuses[len(v.Statuses)-1].Type)
		for _, v := range v.List.Items {
			fmt.Printf("\t%v %v %v %v %v\n", v.ProductID, v.Name, v.Status, v.QuantityRequested, v.QuantityReceived)
		}
	}
}
