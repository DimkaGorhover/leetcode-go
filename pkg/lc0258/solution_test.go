package lc0258

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addDigits(t *testing.T) {
	t.Parallel()
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `test_001`,
			args: args{
				num: 38,
			},
			want: 2,
		},
		{
			name: `test_002`,
			args: args{
				num: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addDigits(tt.args.num))
		})
	}
}
