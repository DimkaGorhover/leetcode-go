package p0008

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: " "}, want: 0},
		{args: args{s: "42"}, want: 42},
		{args: args{s: "    -42"}, want: -42},
		{args: args{s: "words and 987"}, want: 0},
		{args: args{s: "987  "}, want: 987},
		{args: args{s: "-91283472332"}, want: math.MinInt32},
		{args: args{s: "1291283472332"}, want: math.MaxInt32},
		{args: args{s: "4193 with words"}, want: 4_193},
	}
	/*

	 */
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, myAtoi(tt.args.s))
		})
	}
}
