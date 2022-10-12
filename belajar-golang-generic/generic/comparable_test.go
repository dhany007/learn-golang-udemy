package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1 T, value2 T) bool {
	return value1 == value2
}

func TestComparable(t *testing.T) {
	assert.Equal(t, IsSame("andi", "andi"), true)
	assert.Equal(t, IsSame("andi", "budi"), false)
	assert.Equal(t, IsSame(1, 2), false)
	assert.Equal(t, IsSame(1, 1), true)
	assert.Equal(t, IsSame(true, false), false)
	assert.Equal(t, IsSame(true, true), true)
}
