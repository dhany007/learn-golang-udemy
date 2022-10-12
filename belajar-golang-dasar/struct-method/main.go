package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hallo", name, "my name is", customer.Name)
}

func main() {
	fmt.Println("struct method")

	rully := Customer{Name: "Rully"}
	rully.sayHello("dhany")
}
