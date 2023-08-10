package p0146

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Parallel()

	t.Run(`test_001`, func(t *testing.T) {

		cache := Constructor(2)

		cache.Put(2, 1)
		cache.Put(2, 2)

		assert.Equal(t, 2, cache.Get(2))

		cache.Put(1, 1)
		cache.Put(4, 1)

		assert.Equal(t, -1, cache.Get(2))
	})

	t.Run(`test_002`, func(t *testing.T) {

		cache := Constructor(2)

		cache.Put(1, 1)
		cache.Put(2, 1)

		assert.Equal(t, 1, cache.Get(1))

		cache.Put(3, 3)

		assert.Equal(t, -1, cache.Get(2))

		cache.Put(4, 4)

		assert.Equal(t, -1, cache.Get(1))
		assert.Equal(t, 3, cache.Get(3))
		assert.Equal(t, 4, cache.Get(4))
	})
}
