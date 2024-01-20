package lc0303

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		a := Constructor([]int{-2, 0, 3, -5, 2, -1})
		assert.Equal(t, 1, a.SumRange(0, 2))
		assert.Equal(t, -1, a.SumRange(2, 5))
		assert.Equal(t, -3, a.SumRange(0, 5))
	})
}
