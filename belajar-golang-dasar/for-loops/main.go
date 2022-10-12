package main

import "fmt"

func main() {
	fmt.Println("for loops")

	counter := 1
	for counter <= 10 {
		fmt.Println("Counter: ", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i : ", i)
	}

	tempSlice := []string{"dhany", "putra", "aritonang"}
	for i := 0; i < len(tempSlice); i++ {
		fmt.Println(tempSlice[i])
	}

	for index, name := range tempSlice {
		fmt.Println("index : ", index, "name : ", name)
	}
}
