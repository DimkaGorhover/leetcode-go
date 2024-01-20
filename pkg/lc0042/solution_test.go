package lc0042

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trap(t *testing.T) {
	t.Parallel()
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `[0,1,0,2,1,0,1,3,2,1,2,1] ==> 6`,
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: `[4,2,0,3,2,5] ==> 9`,
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
