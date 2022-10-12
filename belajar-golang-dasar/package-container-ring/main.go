package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("circular list")

	// create circular list 5 tempat
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
