package dbs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Car struct {
	Make  *string `json:"make"`
	Model *string `json:"model"`
	Color *string `json:"color"`
	Power *string `json:"power"`
	Year  *string `json:"year"`
	id    int8
}

var db = nil

func init() {
	connStr := "postgres://postgres:pgadmpwd160923@localhost/go-postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not establish connection to database")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to database failed", err)
	}

	fmt.Println("Connection to database established")
}

func runRead(query string) {
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error while fetching data")
		log.Fatal(err)
	}

	return rows
}

func runWrite(query string) {

}
