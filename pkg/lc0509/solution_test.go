package lc0509

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `test_001`,
			args: args{n: 2},
			want: 1,
		},
		{
			name: `test_002`,
			args: args{n: 3},
			want: 2,
		},
		{
			name: `test_003`,
			args: args{n: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fib(tt.args.n))
		})
	}
}
