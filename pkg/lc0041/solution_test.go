package lc0041

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
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
			name: `[1,2,0]=3`,
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		},
		{
			name: `[3,4,-1,1]=2`,
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: `[7,8,9,11,12]=1`,
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
		{
			name: `[0]=1`,
			args: args{
				nums: []int{0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
