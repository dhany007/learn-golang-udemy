package main

import "fmt"

func main()  {
	var nama = "nama"
	fmt.Println(nama)

	nama = "saya"
	fmt.Println(nama)

	// tanpa var
	age := 30
	fmt.Println(age)

	age = 40
	fmt.Println(age)

	//multi variable
	var (
		firstname = "dhany"
		lastname = "aritonang"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

	// konstan
	const phi = 3.14

	fmt.Println(phi)

	const (
		a = "A"
		b = "B"
	)
	fmt.Println(a)
	fmt.Println(b)

}