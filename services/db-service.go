package DatabaseService

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

var db = nil //error here - use of untyped nil in variable declaration
//how to globally store the db object once the connection is made and share it in different functions

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
	} else {
		fmt.Println("Connection to database established")
	}
}

func read(query string) (*sql.Rows, error) {
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error while fetching data")
		return nil, err
	}

	return rows, nil
}

func write(make string, model string, color string, power string, year string) ([]Car, error) { //how to make this accept an object of struct Car instead of individual arguments
	cars, err := db.QueryRow(`INSERT INTO public."Cars"("Make", "Model", "Color", "Power", "Year") VALUES ($1, $2, $3, $4, $5) RETURNING *`, make, model, color, power, year).Scan(&cars)
	if err != nil {
		fmt.Println("Error while writing data")
		return nil, err
	}

	return cars, nil
}

func delete(id int8) (string, error) {
	err := db.Query(`DELETE * from public."Cars" where id = $1`, id)
	if err != nil {
		return "Error while deleting data", err
	}

	return "Deleted successfully", nil
}

func closeDB() {
	err := db.Close()
	if err != nil {
		fmt.Println("Connection could not be closed", err)
	}

	fmt.Println("Connection closed")
}
