package lc2336

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, actual, expected int) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("PopSmallest() = %v, want %v", actual, expected)
	}
}

func TestConstructor_001(t *testing.T) {
	t.Parallel()
	set := Constructor()
	assert(t, set.PopSmallest(), 1)
	set.AddBack(1)
	assert(t, set.PopSmallest(), 1)
	assert(t, set.PopSmallest(), 2)
	assert(t, set.PopSmallest(), 3)
	set.AddBack(2)
	set.AddBack(3)
	assert(t, set.PopSmallest(), 2)
	assert(t, set.PopSmallest(), 3)
}

func TestConstructor_002(t *testing.T) {
	t.Parallel()
	set := Constructor()
	assert(t, set.PopSmallest(), 1)
	assert(t, set.PopSmallest(), 2)
	set.AddBack(3)
	assert(t, set.PopSmallest(), 3)
	set.AddBack(2)
	assert(t, set.PopSmallest(), 2)
	assert(t, set.PopSmallest(), 4)
}
