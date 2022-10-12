package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type AnotherSample struct {
	Name string
	Age  int
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()

			if value == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("package reflection")

	sample := Sample{"eko"}
	sampleType := reflect.TypeOf(sample)
	structType := sampleType.Field(0)
	required := structType.Tag.Get("required")
	max := structType.Tag.Get("max")

	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(required)
	fmt.Println(max)

	isValid := IsValid(sample)
	fmt.Println(isValid)

	// tidak ada tag required jadi otomatis true
	sampleLagi := AnotherSample{"", 10}
	isValidLagi := IsValid(sampleLagi)
	fmt.Println(isValidLagi)
}
