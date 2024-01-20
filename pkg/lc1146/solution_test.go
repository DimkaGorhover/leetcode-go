package lc1146

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnapshotArray_001(t *testing.T) {
	t.Parallel()

	snapshotArr := Constructor(3)
	snapshotArr.Set(0, 5)
	assert.Equal(t, 0, snapshotArr.Snap())
	snapshotArr.Set(0, 6)
	assert.Equal(t, 5, snapshotArr.Get(0, 0))
}

func TestSnapshotArray_002(t *testing.T) {
	t.Parallel()

	snapshotArr := Constructor(1)
	snapshotArr.Set(0, 4)
	snapshotArr.Set(0, 16)
	snapshotArr.Set(0, 13)
	assert.Equal(t, 0, snapshotArr.Snap())
	assert.Equal(t, 13, snapshotArr.Get(0, 0))
	assert.Equal(t, 1, snapshotArr.Snap())
}

func TestSnapshotArray_003(t *testing.T) {
	t.Parallel()

	snapshotArr := Constructor(2)
	assert.Equal(t, 0, snapshotArr.Snap())
	assert.Equal(t, 0, snapshotArr.Get(1, 0))
	assert.Equal(t, 0, snapshotArr.Get(0, 0))
	snapshotArr.Set(1, 8)
	assert.Equal(t, 0, snapshotArr.Get(1, 0))
	snapshotArr.Set(0, 20)
	assert.Equal(t, 0, snapshotArr.Get(0, 0))
	snapshotArr.Set(0, 7)
}

func TestSnapshotArray_004(t *testing.T) {
	t.Parallel()

	snapshotArr := Constructor(3)
	snapshotArr.Set(1, 18)
	snapshotArr.Set(1, 4)
	assert.Equal(t, 0, snapshotArr.Snap())
	assert.Equal(t, 0, snapshotArr.Get(0, 0))
	snapshotArr.Set(0, 20)
	assert.Equal(t, 1, snapshotArr.Snap())
	snapshotArr.Set(0, 2)
	snapshotArr.Set(1, 1)
	assert.Equal(t, 4, snapshotArr.Get(1, 1))
	assert.Equal(t, 4, snapshotArr.Get(1, 0))
}
