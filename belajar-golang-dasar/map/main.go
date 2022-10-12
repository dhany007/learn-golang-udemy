package main

import "fmt"

func main() {
	fmt.Println("hello")

	person := map[string]string{
		"name":   "dhany",
		"alamat": "simanabun",
	}
	fmt.Println(person)
	fmt.Println(person["name"])

	person["title"] = "Mr. "
	fmt.Println(person)

	book := make(map[string]string)
	book["name"] = "Buku"
	book["Year"] = "2021"

	fmt.Println(book)
	delete(book, "year")
	fmt.Println(book)
}
