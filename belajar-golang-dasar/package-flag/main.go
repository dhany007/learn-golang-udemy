package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("flag")

	host := flag.String("host", "localhost", "put your host")
	user := flag.String("user", "root", "put your user")
	password := flag.String("password", "root", "put your password")

	flag.Parse()

	fmt.Println("host =", *host)
	fmt.Println("user =", *user)
	fmt.Println("password =", *password)

	// how to run
	// go run main.go -host=localhost -user=root -password=root
}
