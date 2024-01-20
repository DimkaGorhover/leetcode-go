package lc0919

import (
	"reflect"
	"testing"
)

func Test_newQueue(t *testing.T) {
	t.Parallel()

	t.Run(`test_001`, func(t *testing.T) {

		assertEquals := func(expected, actual any) {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf(`expected = %v; actual = %v`, expected, actual)
			}
		}

		q := newQueue[int]()
		assertEquals(0, q.Size())
		assertEquals([]int{}, q.Values())

		q.Push(10)
		assertEquals(10, q.Peek())
		assertEquals([]int{10}, q.Values())
		assertEquals(1, q.Size())

		assertEquals(10, q.Peek())
		assertEquals([]int{10}, q.Values())
		assertEquals(1, q.Size())

		assertEquals(10, q.Peek())
		assertEquals([]int{10}, q.Values())
		assertEquals(1, q.Size())

		assertEquals(10, q.Poll())
		assertEquals([]int{}, q.Values())
		assertEquals(0, q.Size())

		q.PushAll(1, 2, 3, 4, 5)
		assertEquals([]int{1, 2, 3, 4, 5}, q.Values())
		assertEquals(1, q.Poll())
		assertEquals([]int{2, 3, 4, 5}, q.Values())
	})
}
