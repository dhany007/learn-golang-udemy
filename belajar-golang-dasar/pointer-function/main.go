package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "USA"
}

func ChangeAddressToIndonesiaPointer(address *Address) {
	address.Country = "USA"
}

func main() {
	fmt.Println("pointer di function")

	address := Address{"Malang", "Jakarta", "Indonesia"}

	fmt.Println("address", address)   //{Malang Jakarta Indonesia}
	ChangeAddressToIndonesia(address) // tidak berubah

	fmt.Println("address", address) // {Malang Jakarta Indonesia}

	// pake pointer
	fmt.Println("address", address)           //{Malang Jakarta Indonesia}
	ChangeAddressToIndonesiaPointer(&address) // berubah

	fmt.Println("address", address) // {Malang Jakarta USA}

	// kalau data struct nya besar, diusahakan pake pointer daripada pake value
}
