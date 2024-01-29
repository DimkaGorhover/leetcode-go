package lc0232

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		q := Constructor()
		q.Push(1)
		q.Push(2)
		assert.Equal(t, 1, q.Peek())
		assert.Equal(t, 1, q.Pop())
		assert.False(t, q.Empty())
		assert.Equal(t, 2, q.Pop())
		assert.True(t, q.Empty())
	})
}
