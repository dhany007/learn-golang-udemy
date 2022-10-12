package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "01",
		Name:     "Dell",
		ImageURL: "http://www.example.com/image.jpg",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonRequest := `{"id":"01","name":"Dell","image_url":"http://www.example.com/image.jpg"}`
	jsonBytes := []byte(jsonRequest)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
