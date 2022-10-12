package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	costumer := Customer{
		FirstName:  "Dhany",
		MiddleName: "Putra",
		LastName:   "Aritonang",
		Age:        20,
		Married:    false,
	}

	bytes, _ := json.Marshal(costumer)

	fmt.Println(string(bytes))
}
