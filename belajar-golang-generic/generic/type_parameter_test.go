package generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type constraint
func Length[T interface{}]() {

}

// type any itu interface kosong
// type constraint
func LengthAny[T any](param T) T {
	fmt.Println("param", param)
	return param
}

func TestSample(t *testing.T) {
	// resultString := LengthAny[string]("Dhany") gini juga bisa
	resultString := LengthAny("Dhany")
	assert.Equal(t, "Dhany", resultString)

	// resultNumber := LengthAny[int](3) gini juga bisa
	resultNumber := LengthAny(3)
	assert.Equal(t, 3, resultNumber)
}
