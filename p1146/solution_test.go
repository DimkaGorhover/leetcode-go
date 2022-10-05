package p1146

import (
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v; expected = %v", actual, expected)
	}
}

func TestSnapshotArray_001(t *testing.T) {
	snapshotArr := Constructor(3)
	snapshotArr.Set(0, 5)
	assertEquals(t, 0, snapshotArr.Snap())
	snapshotArr.Set(0, 6)
	assertEquals(t, 5, snapshotArr.Get(0, 0))
}

func TestSnapshotArray_002(t *testing.T) {
	snapshotArr := Constructor(1)
	snapshotArr.Set(0, 4)
	snapshotArr.Set(0, 16)
	snapshotArr.Set(0, 13)
	assertEquals(t, 0, snapshotArr.Snap())
	assertEquals(t, 13, snapshotArr.Get(0, 0))
	assertEquals(t, 1, snapshotArr.Snap())
}

func TestSnapshotArray_003(t *testing.T) {
	snapshotArr := Constructor(2)
	assertEquals(t, 0, snapshotArr.Snap())
	assertEquals(t, 0, snapshotArr.Get(1, 0))
	assertEquals(t, 0, snapshotArr.Get(0, 0))
	snapshotArr.Set(1, 8)
	assertEquals(t, 0, snapshotArr.Get(1, 0))
	snapshotArr.Set(0, 20)
	assertEquals(t, 0, snapshotArr.Get(0, 0))
	snapshotArr.Set(0, 7)
}

func TestSnapshotArray_004(t *testing.T) {
	snapshotArr := Constructor(3)
	snapshotArr.Set(1, 18)
	snapshotArr.Set(1, 4)
	assertEquals(t, 0, snapshotArr.Snap())
	assertEquals(t, 0, snapshotArr.Get(0, 0))
	snapshotArr.Set(0, 20)
	assertEquals(t, 1, snapshotArr.Snap())
	snapshotArr.Set(0, 2)
	snapshotArr.Set(1, 1)
	assertEquals(t, 4, snapshotArr.Get(1, 1))
	assertEquals(t, 4, snapshotArr.Get(1, 0))
}
