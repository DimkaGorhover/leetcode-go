package lc2108

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstPalindrome(t *testing.T) {
	t.Parallel()
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `test_001`,
			args: args{
				words: []string{"abc", "car", "ada", "racecar", "cool"},
			},
			want: `ada`,
		},
		{
			name: `test_002`,
			args: args{
				words: []string{"notapalindrome", "racecar"},
			},
			want: `racecar`,
		},
		{
			name: `test_003`,
			args: args{
				words: []string{"def", "ghi"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, firstPalindrome(tt.args.words))
		})
	}
}
