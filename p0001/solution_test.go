package p0001

import (
	"reflect"
	"testing"
)

type testcase struct {
	input    []int
	target   int
	expected []int
}

func TestAverage(t *testing.T) {

	testsCases := []testcase{
		{[]int{7, 11, 15, 2}, 9, []int{0, 3}},
	}

	for _, test := range testsCases {
		actual := twoSum(test.input, test.target)
		if !(reflect.DeepEqual(test.expected, actual)) {
			t.Error(
				"For", test.input,
				"expected", test.expected,
				"got", actual,
			)
		}
	}
}
