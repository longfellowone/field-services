package main

import (
	"database/sql"
	"field/pkg/ordering"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	defaultGRPCPort = 9090
	defaultDBHost   = "localhost"
	defaultDBPort   = 5432
	defaultDBName   = "default"
	defaultDBUser   = "default"
	defaultDBPasswd = "password"
	sslMode         = "disable"
	inMemory        = true
)

func main() {

	// flags := flag.Parse()

	var (
		dbHost                   = defaultDBHost
		dbPort                   = defaultDBPort
		dbUser                   = defaultDBUser
		dbPasswd                 = defaultDBPasswd
		dbName                   = defaultDBName
		postgresConnectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPasswd, dbName, sslMode)
	)

	var fs *ordering.Service

	if inMemory {
		fs = initializeFieldServicesInMemory()
	} else {
		db, err := sql.Open("postgres", postgresConnectionString)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}

		fs = initializeFieldServices(db)
	}

	result, _ := fs.FindAllOrders()

	fmt.Printf("Result: %s", result)
}
