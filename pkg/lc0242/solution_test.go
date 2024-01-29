package lc0242

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `test_001`,
			args: args{
				s: `anagram`,
				t: `nagaram`,
			},
			want: true,
		},
		{
			name: `test_002`,
			args: args{
				s: `rat`,
				t: `car`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAnagram(tt.args.s, tt.args.t))
		})
	}
}
