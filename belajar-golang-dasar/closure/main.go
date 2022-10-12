package main

import "fmt"

func main() {
	fmt.Println("closure")
	counter := 0
	name := "dhany"

	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
