package lc0076

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: ``, args: args{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
		{name: ``, args: args{s: "a", t: "a"}, want: "a"},
		{name: ``, args: args{s: "cabwefgewcwaefgcf", t: "cae"}, want: "cwae"},
		{name: ``, args: args{s: "a", t: "b"}, want: ""},
		{name: ``, args: args{s: "a", t: "aa"}, want: ""},
		{name: ``, args: args{s: "bba", t: "ab"}, want: "ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minWindow(tt.args.s, tt.args.t))
		})
	}
}
