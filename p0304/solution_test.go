package p0304

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		m := Constructor([][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		})

		fmt.Printf("%s", m)

		assert.Equal(t, 8, m.SumRegion(2, 1, 4, 3))
		assert.Equal(t, 11, m.SumRegion(1, 1, 2, 2))
		assert.Equal(t, 12, m.SumRegion(1, 2, 2, 4))
	})

}
