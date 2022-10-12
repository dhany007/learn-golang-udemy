package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	costumer := Customer{
		FirstName:  "Dhany",
		MiddleName: "Putra",
		LastName:   "Aritonang",
		Age:        20,
		Married:    false,
		Hobbies:    []string{"Game", "Baca"},
	}

	bytes, _ := json.Marshal(costumer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Dhany","MiddleName":"Putra","LastName":"Aritonang","Age":20,"Married":false,"Hobbies":["Game","Baca"]}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	costumer := Customer{
		FirstName: "Dhany",
		Addresses: []Address{
			{
				Street:     "Jalan 1",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
			{
				Street:     "Jalan 2",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
		},
	}

	bytes, _ := json.Marshal(costumer)

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Dhany","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"12345"}]}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"12345"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)

	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan 1",
			Country:    "Indonesia",
			PostalCode: "12345",
		},
		{
			Street:     "Jalan 2",
			Country:    "Indonesia",
			PostalCode: "12345",
		},
	}

	bytes, _ := json.Marshal(addresses)

	fmt.Println(string(bytes))
}
