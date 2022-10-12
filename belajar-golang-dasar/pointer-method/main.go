package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fmt.Println("pointer di method")

	// saangat direkomendasikan membuat pointer di method struct,
	// supaya tidak duplikasi2 data

	dhany := Man{"Dhany"}
	fmt.Println(dhany.Name)
	dhany.Married() // name tidak berubah
	fmt.Println(dhany.Name)

	eko := Man{"Eko"}
	fmt.Println(eko.Name)
	eko.MarriedPointer() // name berubah karena ada pointer
	fmt.Println(eko.Name)

	// disarankan kalau pake struct, pake pointer untuk methodnya
}
