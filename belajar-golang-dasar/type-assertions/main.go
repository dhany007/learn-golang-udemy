package main

import (
	"fmt"
)

func Random() interface{} {
	// return true
	// return "OK"
	return 100
}

func main() {
	fmt.Println("type assertions")
	// merubah tipe data menjadi tipe data yang kita inginkan

	result := Random()

	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt) // panic atau error

	// menggunakan switch expressions untuk type assertions

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Int")
	default:
		fmt.Println("Unknown Type")
	}
}
