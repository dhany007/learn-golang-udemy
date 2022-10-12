package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONObjectDecode(t *testing.T) {
	jsonRequest := `{"FirstName": "kalai", "MiddleName": "dhany", "LastName": "saragih", "Age" : 19, "Married": false}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
