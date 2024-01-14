package p0284

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {

	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		i := Constructor(NewIterator([]int{1, 2, 3, 4, 5}))
		assert.True(t, i.hasNext())
		assert.Equal(t, 1, i.peek())
		assert.Equal(t, 1, i.peek())
		assert.Equal(t, 1, i.peek())
		assert.Equal(t, 1, i.next())
		assert.Equal(t, 2, i.next())
		assert.True(t, i.hasNext())
		assert.Equal(t, 3, i.next())
		assert.True(t, i.hasNext())
		assert.Equal(t, 4, i.next())
		assert.True(t, i.hasNext())
		assert.Equal(t, 5, i.peek())
		assert.True(t, i.hasNext())
		assert.Equal(t, 5, i.peek())
		assert.True(t, i.hasNext())
		assert.Equal(t, 5, i.next())
		assert.False(t, i.hasNext())
	})

}
