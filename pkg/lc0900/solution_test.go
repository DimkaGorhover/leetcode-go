package lc0900

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRLEIterator_Next(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {

		itr := Constructor([]int{3, 8, 0, 9, 2, 5})

		assert.Equal(t, 8, itr.Next(2))
		assert.Equal(t, 8, itr.Next(1))
		assert.Equal(t, 5, itr.Next(1))
		assert.Equal(t, -1, itr.Next(2))
	})
}
