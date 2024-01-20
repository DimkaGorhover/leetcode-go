package lc0128

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

//go:embed test_004.txt
var test004 string

func Test_longestConsecutive(t *testing.T) {
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
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			name: `test_002`,
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
		{
			name: `test_003`,
			args: args{
				nums: []int{100, 1, 200, 3, 2, 4},
			},
			want: 4,
		},
		{
			name: `test_004`,
			args: args{
				nums: func() []int {
					arr := strings.Split(test004[1:len(test004)-1], ",")
					nums := make([]int, len(arr))
					for i := range arr {
						num, _ := strconv.Atoi(arr[i])
						nums[i] = num
					}
					return nums
				}(),
			},
			want: 100_000,
		},
		{
			name: `test_005`,
			args: args{
				nums: []int{1, 0, -1},
			},
			want: 3,
		},
		{
			name: `test_006`,
			args: args{
				nums: []int{0, 0, 1, -1},
			},
			want: 3,
		},
		{
			name: `test_007`,
			args: args{
				nums: []int{9, -8, 9, 8, -7, 9, -4, 6, 5, 5, 6, 7, -9, -5, -4, 6, -8, -1, 8, 0, 1, 5, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestConsecutive(tt.args.nums))
		})
	}
}
