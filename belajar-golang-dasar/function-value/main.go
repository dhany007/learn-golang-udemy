package main

import "fmt"

func sayGoodBye(nama string) string {
	return "Goodbye " + nama
}

func main() {
	fmt.Println("function value")

	goodbye := sayGoodBye

	result := goodbye("eko")

	fmt.Println(result)
}
