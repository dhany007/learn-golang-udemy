package main

import "fmt"

func main() {
	fmt.Println("if expression")

	var name = "eko"

	fmt.Println(name)

	if name == "dhany" {
		fmt.Println("Hello dhany")
	} else if name == "eko" {
		fmt.Println("nama kamu eko ?")
	} else {
		fmt.Println("Hallo, nama kamu siapa ?")
	}

	if length := len(name); length > 20 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
