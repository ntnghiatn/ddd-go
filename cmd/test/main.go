package main

import (
	json2 "encoding/json"
	"github.com/ntnghiatn/ddd-go/aggregate"
	"log"
)

func main() {

	log.Println("-----TEST-----")
	log.Println("NewCustomer")

	customer, err := aggregate.NewCustomer("TEST")
	if err != nil {
		panic("Err create Customer!!!")
	}
	json, err := json2.Marshal(customer)
	if err != nil {
		panic("Err create Customer!!!")
	}
	log.Println(json)
}
