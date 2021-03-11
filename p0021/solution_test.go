package p0021

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {

	list1 := ListNodeOf(1, 3, 5)
	list2 := ListNodeOf(2, 4, 6)

	expected := ListNodeOf(1, 2, 3, 4, 5, 6)
	actual := mergeTwoLists(list1, list2)

	if !(reflect.DeepEqual(expected, actual)) {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}
