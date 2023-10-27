package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"services/DatabaseService"
)

type Response struct {
	message string
	data    *DatabaseService.Car
}

func getCars() DatabaseService.Car {
	rows, err := DatabaseService.read("select * from public.\"Cars\"")
	if err != nil {
		log.Fatal(err)
	}

	var cars []DatabaseService.Car
	for rows.Next() {
		var car DatabaseService.Car
		if err := rows.Scan(&car.Make, &car.Model, &car.Color, &car.Power, &car.Year, &car.id); err != nil { //how to read as a single car object instead of individual fields
			log.Fatal(err)
		}
		cars = append(cars, car)
	}

	return cars
}

func addCar(r *http.Request) []DatabaseService.Car {
	decoder := json.NewDecoder(r.Body)
	var car DatabaseService.Car
	err := decoder.Decode(&car)
	if err != nil {
		fmt.Println("Error while reading request body")
		log.Fatal(err)
	}

	var cars []DatabaseService.Car
	cars, err := DatabaseService.write(car.Make, car.Model, car.Color, car.Power, car.Year)
	if err != nil {
		log.Fatal(err)
	}

	return cars
}

func removeCar(r *http.Request) Response {
	decoder := json.NewDecoder(r.Body)
	var car DatabaseService.Car
	err := decoder.Decode(&car.id)
	if err != nil {
		fmt.Println("Error while reading request body")
		log.Fatal(err)
	}

	//return json with key message and value as the returned value from db function
	msg, err := DatabaseService.delete(car.id)
	if err != nil {
		log.Fatal(err)
	}

	resp := Response{message: msg}
	return resp
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cars" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		cars := getCars()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	case "POST":
		cars := addCar(r)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	case "DELETE":
		resp := removeCar()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "406 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/cars", routeHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	} else {
		DatabaseService.init()
	}
}
