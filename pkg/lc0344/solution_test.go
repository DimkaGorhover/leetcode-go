package lc0344

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	t.Parallel()
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: `test_001`,
			args: args{
				s: []byte("hello world"),
			},
			want: []byte("dlrow olleh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			assert.Equal(t, tt.want, tt.args.s)
		})
	}
}
