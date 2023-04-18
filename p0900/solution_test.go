package p0900

import (
	"reflect"
	"testing"
)

func TestRLEIterator_Next(t *testing.T) {
	t.Parallel()

	assertEquals := func(expected, actual int) {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Next() = %v, want %v", actual, expected)
		}
	}

	itr := Constructor([]int{3, 8, 0, 9, 2, 5})

	assertEquals(8, itr.Next(2))
	assertEquals(8, itr.Next(1))
	assertEquals(5, itr.Next(1))
	assertEquals(-1, itr.Next(2))

}
