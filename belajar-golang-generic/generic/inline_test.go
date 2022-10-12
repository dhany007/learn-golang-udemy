package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface {
	float32 | float64 | int | int64
}](value1 T, value2 T) T {
	min := value1

	if value2 < value1 {
		min = value2
	}
	return min
}

func TestInlineTest(t *testing.T) {
	assert.Equal(t, 1, FindMin(1, 2))
	assert.Equal(t, 1.0, FindMin(2.0, 1.0))
	assert.Equal(t, int64(3), FindMin(int64(3), int64(10)))
}

// [T []E, E any] artinya T harus diikuti slice E dan E boleh tipe apapun
func GetFirst[T []E, E any](data T) E {
	return data[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{"Dhany", "Kalai", "Aritonang"}
	assert.Equal(t, "Dhany", GetFirst(names))

	numbers := []int{1, 2, 3}
	assert.Equal(t, 1, GetFirst(numbers))
}
