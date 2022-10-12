package generic

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, b := range bag {
		fmt.Println(b)
	}
}

func TestTypeGeneric(t *testing.T) {
	names := Bag[string]{"Dhany", "Budi"}
	PrintBag(names)
	intNumbers := Bag[int]{1, 2, 3}
	PrintBag(intNumbers)
	floatNumbers := Bag[float64]{1, 2, 3}
	PrintBag(floatNumbers)
}
