package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Kalai",
		MiddleName: "Dhany",
		LastName:   "Saragih",
	}

	encoder.Encode(customer)

	fmt.Println(customer)
}
