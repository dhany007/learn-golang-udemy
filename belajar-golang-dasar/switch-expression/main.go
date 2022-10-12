package main

import "fmt"

func main() {
	fmt.Println("switch expression")

	var name = "bahabahbahabhbshdbahsbdhabsdhabd"

	switch name {
	case "eko":
		fmt.Println("kamu adalah eko")
	default:
		fmt.Println("kamu bukan eko")
	}

	switch length := len(name); length > 20 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 20:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
