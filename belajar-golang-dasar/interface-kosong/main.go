package main

import "fmt"

func Up(number int) interface{} {
	if number == 1 {
		return 1
	} else if number == 2 {
		return 2
	} else {
		return number
	}
}

func main() {
	fmt.Println("Interface kosong")

	var data interface{} = Up(10)

	fmt.Println(data)
}
