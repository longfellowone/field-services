package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/getting-started/bookshelf/db_mysql.go

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func Connect(cs Config) (*sql.DB, error) {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cs.DBUser, cs.DBPassword, cs.DBHost, cs.DBPort, cs.DBName)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
