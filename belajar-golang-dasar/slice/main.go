package main

import "fmt"

func main() {
	fmt.Println(("Hallo"))

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Nopember",
		"Desember",
	}

	fmt.Println(months)

	var slice1 = months[4:7]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1[0] = "Mei Update"

	fmt.Println(months)
}
