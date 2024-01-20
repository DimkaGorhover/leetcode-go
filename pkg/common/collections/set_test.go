package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		aSet := NewSet[int]()

		assert.True(t, aSet.AddAll(1, 1, 2, 3, 1, 2, 3))
		assert.Equal(t, 3, aSet.Size())
		assert.Equal(t, []int{1, 2, 3}, aSet.Values())
		assert.Equal(t, `[1, 2, 3]`, aSet.String())
		assert.True(t, aSet.Contains(1))
		assert.True(t, aSet.Contains(2))
		assert.True(t, aSet.Contains(3))
		assert.False(t, aSet.Contains(4))

		assert.False(t, aSet.AddAll(1, 1, 2, 3))
		assert.Equal(t, 3, aSet.Size())
		assert.Equal(t, []int{1, 2, 3}, aSet.Values())
		assert.Equal(t, `[1, 2, 3]`, aSet.String())
		assert.True(t, aSet.Contains(1))
		assert.True(t, aSet.Contains(2))
		assert.True(t, aSet.Contains(3))
		assert.False(t, aSet.Contains(4))

		assert.True(t, aSet.Add(4))
		assert.Equal(t, 4, aSet.Size())
		assert.Equal(t, []int{1, 2, 3, 4}, aSet.Values())
		assert.Equal(t, `[1, 2, 3, 4]`, aSet.String())
		assert.True(t, aSet.Contains(1))
		assert.True(t, aSet.Contains(2))
		assert.True(t, aSet.Contains(3))
		assert.True(t, aSet.Contains(4))

		assert.True(t, aSet.Remove(4))
		assert.Equal(t, []int{1, 2, 3}, aSet.Values())
		assert.Equal(t, `[1, 2, 3]`, aSet.String())
		assert.Equal(t, 3, aSet.Size())
		assert.True(t, aSet.Contains(1))
		assert.True(t, aSet.Contains(2))
		assert.True(t, aSet.Contains(3))
		assert.False(t, aSet.Contains(4))

		assert.False(t, aSet.Remove(4))
		assert.Equal(t, 3, aSet.Size())
		assert.Equal(t, []int{1, 2, 3}, aSet.Values())
		assert.Equal(t, `[1, 2, 3]`, aSet.String())
	})
}
