package lc3005

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxFrequencyElements(t *testing.T) {
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
				nums: []int{1, 2, 2, 3, 1, 4},
			},
			want: 4,
		},
		{
			name: `test_002`,
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: `test_003`,
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: `test_004`,
			args: args{
				nums: []int{10},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxFrequencyElements(tt.args.nums))
		})
	}
}
