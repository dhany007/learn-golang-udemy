package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("package regex")
	regex := regexp.MustCompile(`a([a-z])d`)

	fmt.Println(regex.MatchString("andi"))
	fmt.Println(regex.MatchString("ndi"))
}
