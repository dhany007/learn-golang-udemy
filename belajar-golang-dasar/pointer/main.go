package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	fmt.Println("pointer")
	// go lang pass by value
	// variabel yang dikirim ke func lain merupakan duplikat dari value nya

	// //ex pass by value
	// address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	// address2 := address1

	// // city addres1 tidak akan mengubah nilai city addres2
	// address1.City = "Jakarta Selatan"

	// fmt.Println(address1) // {Jakarta Selatan Sumatera Utara Indonesia}
	// fmt.Println(address2) // {Medan Sumatera Utara Indonesia}

	// // pointer = pass by reference
	// // kemampuan untuk membuat kelokasi reference sama, tanpa duplikasi data
	// // '&' : pointer

	// address3 := &address1

	// // address 1 akan mengubah address 3
	// address1.City = "Jakarta Barat"

	// fmt.Println(address1) // {Jakarta Barat Sumatera Utara Indonesia}
	// fmt.Println(address3) // &{Jakarta Barat Sumatera Utara Indonesia}

	// atau pake variabel
	var address4 Address = Address{"Medan", "Sumut", "Indo"}
	var address5 *Address = &address4
	var address6 *Address = &address4
	address5.City = "Bandung"
	fmt.Println("address4", address4) //{Bandung Sumut Indo}
	fmt.Println("address5", address5) //&{Bandung Sumut Indo}

	// operator *
	// // address5 tidak akan mengubah isi address4
	// // address5 akan membuat data baru
	// address5 = &Address{"simanabun", "siloukahean", "simalungun"}
	// fmt.Println("address4", address4) //{Bandung Sumut Indo}
	// fmt.Println("address5", address5) // &{simanabun siloukahean simalungun}

	// bagaimana jika kalau kita mengubah juga address4 dan memaksa ?
	// pake operator *
	*address5 = Address{"simanabun2", "siloukahean2", "simalungun2"}
	fmt.Println("address4", address4) //{simanabun2 siloukahean2 simalungun2}
	fmt.Println("address5", address5) // &{simanabun siloukahean simalungun}
	fmt.Println("address6", address6) // &{simanabun siloukahean simalungun}

	// membuat pointer dengan data kosong (new)

	alamatBaru := new(Address)
	fmt.Println(alamatBaru)
	alamatBaru.City = "Jakarta"
	fmt.Println(alamatBaru)
}
