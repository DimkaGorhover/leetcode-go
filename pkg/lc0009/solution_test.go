package lc0009

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	t.Parallel()
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{121},
			want: true,
		},
		{
			args: args{12345678987654321},
			want: true,
		},
		{
			args: args{123456789987654321},
			want: true,
		},
		{
			args: args{1234567987654321},
			want: false,
		},
		{
			args: args{-121},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.x))
		})
	}
}
