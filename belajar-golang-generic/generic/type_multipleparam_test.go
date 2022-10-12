package generic

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println("param1", param1)
	fmt.Println("param2", param2)
}

func TestTypeParameter(t *testing.T) {
	MultipleParameter("andi", 2)
	MultipleParameter(2, "budi")
}
