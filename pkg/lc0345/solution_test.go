package lc0345

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `test_001`,
			args: args{
				s: "hello",
			},
			want: "holle",
		},
		{
			name: `test_002`,
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseVowels(tt.args.s))
		})
	}
}
