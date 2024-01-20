package lc1122

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	t.Parallel()
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
				arr2: []int{2, 1, 4, 3, 9, 6},
			},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			args: args{
				arr1: []int{28, 6, 22, 8, 44, 17},
				arr2: []int{22, 28, 8, 6},
			},
			want: []int{22, 28, 8, 6, 17, 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, relativeSortArray(tt.args.arr1, tt.args.arr2))
		})
	}
}
