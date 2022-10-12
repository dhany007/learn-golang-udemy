package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}

func main() {
	fmt.Println("nill")

	var person map[string]string = nil
	fmt.Println(person)

	data := NewMap("")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
