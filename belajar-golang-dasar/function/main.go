package main

import "fmt"

func sayHello() {
	fmt.Println("Hallo")
}

func sayNameHello(name string) {
	fmt.Println("Hallo", name)

}

func getHallo(name string) string {
	result := "hallo " + name
	return result
}

func fullname() (string, string) {
	return "dhany", "aritonang"
}

func getFullname() (string, string, string) {
	return "dhany", "putra", "aritonang"
}

func namedFullname() (firstname string, middlename string, lastname string) {
	firstname = "dhany"
	middlename = "putra"
	lastname = "aritonang"

	return
}

func getSummary(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println("function")
	sayHello()
	sayNameHello("Dhany")

	result := getHallo("eko")
	fmt.Println(result)

	firstname, lastname := fullname()
	fmt.Println("fullname", firstname, lastname)

	myFirstname, _, myLastname := getFullname()
	fmt.Println(myFirstname)
	fmt.Println(myLastname)

	namedFirst, namedMiddle, namedLast := namedFullname()

	fmt.Println(namedFirst)
	fmt.Println(namedMiddle)
	fmt.Println(namedLast)

	myTotal := getSummary(10, 10, 10)

	fmt.Println(myTotal)

	myNumbers := []int{10, 20, 30, 40}
	fmt.Println(getSummary(myNumbers...))
}
