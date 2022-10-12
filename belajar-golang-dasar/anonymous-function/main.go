package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu diblok", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	fmt.Println("Anonymous function")
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("eko", blacklist)
	registerUser("anjing", blacklist)
}
