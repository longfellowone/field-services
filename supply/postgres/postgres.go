package postgres

import (
	_ "github.com/lib/pq"
)

// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/getting-started/bookshelf/db_mysql.go

//func Connect(dbHost string, dbPort int, dbUser, dbPasswd, dbName string) (*sql.DB, error) {
//	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPasswd, dbHost, dbPort, dbName)
//
//	db, err := sql.Open("postgres", conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	if err = db.Ping(); err != nil {
//		log.Fatal(err)
//	}
//	return db, nil
//}

//db, err := postgres.Connect("localhost", 5432, "default", "password", "default")
//if err != nil {
//	log.Fatal(err)
//}
