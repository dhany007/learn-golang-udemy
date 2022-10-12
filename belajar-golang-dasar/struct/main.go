package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	fmt.Println("struct")

	var dhany Customer
	dhany.Name = "Dhany"
	dhany.Address = "Simanabun"
	dhany.Age = 10

	fmt.Println(dhany)

	joko := Customer{
		Name:    "Joko",
		Address: "jakarta",
		Age:     10,
	}

	fmt.Println(joko)

	budi := Customer{"budi", "indonesia", 12}

	fmt.Println(budi)
}
