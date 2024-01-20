package lc2259

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDigit(t *testing.T) {
	t.Parallel()
	type args struct {
		number string
		digit  byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: ``,
			args: args{"123", '3'},
			want: "12",
		},
		{
			name: ``,
			args: args{"1231", '1'},
			want: "231",
		},
		{
			name: ``,
			args: args{"551", '5'},
			want: "51",
		},
		{
			name: ``,
			args: args{"133235", '3'},
			want: "13325",
		},
		{
			name: ``,
			args: args{"73197", '7'},
			want: "7319",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeDigit(tt.args.number, tt.args.digit))
		})
	}
}
