package belajargolanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	fmt.Println("Belajar Golang Goroutines")

	go HelloWorld()
	fmt.Println("Upss")
	time.Sleep(1 * time.Second)
}
