package lc0013

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `test_001`,
			args: args{
				s: `III`,
			},
			want: 3,
		},
		{
			name: `test_002`,
			args: args{
				s: `LVIII`,
			},
			want: 58,
		},
		{
			name: `test_003`,
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: `test_004`,
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: `test_005`,
			args: args{
				s: "DCCIX",
			},
			want: 709,
		},
		{
			name: `test_006`,
			args: args{
				s: `MCMXCIV`,
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, romanToInt(tt.args.s))
		})
	}
}
