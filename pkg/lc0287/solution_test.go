package lc0287

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	t.Parallel()
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `test_001`,
			args: args{
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: `test_002`,
			args: args{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {

		t.Run(fmt.Sprintf("%s/negative_arr", tt.name), func(t *testing.T) {
			nums := tt.args.nums
			newNums := make([]int, len(nums))
			copy(newNums, nums)

			s := negativeArraySolution{}
			assert.Equal(t, tt.want, s.findDuplicate(newNums))
		})

		t.Run(fmt.Sprintf("%s/hash_map", tt.name), func(t *testing.T) {
			s := hashSolution{}
			assert.Equal(t, tt.want, s.findDuplicate(tt.args.nums))
		})

		t.Run(fmt.Sprintf("%s/default", tt.name), func(t *testing.T) {
			assert.Equal(t, tt.want, findDuplicate(tt.args.nums))
		})

	}
}
