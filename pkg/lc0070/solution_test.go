package lc0070

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{2}, want: 2},
		{args: args{3}, want: 3},
		{args: args{5}, want: 8},
		{args: args{8}, want: 34},
		{args: args{19}, want: 6765},
		{args: args{44}, want: 1_134_903_170},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, climbStairs(tt.args.n))
		})
	}
}
