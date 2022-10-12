package generic

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestGenericStruct(t *testing.T) {
	dataNames := Data[string]{First: "Dhany", Second: "Aritonang"}
	fmt.Printf("data names %+v\n", dataNames)

	dataNumbers := Data[int]{First: 1, Second: 2}
	fmt.Printf("data number %+v\n", dataNumbers)
}
