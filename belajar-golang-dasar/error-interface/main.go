package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak dapat dibagi 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	fmt.Println("error interface")
	// interface kontrak untuk membuat error namanya error interface

	// var contoh error = errors.New("ups error")

	// fmt.Println(contoh.Error())

	hasil, err := Pembagi(10, 10)
	// hasil, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("Hasil =", hasil)
	} else {
		fmt.Println("Error =", err.Error())
	}

}
