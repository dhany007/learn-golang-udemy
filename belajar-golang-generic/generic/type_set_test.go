package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](param1 T, param2 T) T {
	tempMin := param1

	if param2 < param1 {
		tempMin = param2
	}

	return tempMin
}

func TestSetMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 3))
	assert.Equal(t, 1.0, Min(1.0, 3.0))
	assert.Equal(t, 1, Min(3, 1))
	assert.Equal(t, 1000000000, Min(1_000_000_000, 10_000_000_000))

	// disini tidak bisa menggunakan Age, walaupun Age itu dideklarasikan sebagai int
	// karena Age tidak termasuk dalam constraint Number
	// itu hal yang berbeda
	// assert.Equal(t, Age(10), Min[Age](Age(10), Age(10)))
}

// semua yang ada tanda ~ , maka tipe dasarnya bisa digunakan
// contohnya pada Number int diatas

func TestTypeApproximation(t *testing.T) {
	assert.Equal(t, Age(10), Min(Age(10), Age(10)))
}
