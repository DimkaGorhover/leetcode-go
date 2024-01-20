package lc2319

import "testing"

func Test_checkXMatrix(t *testing.T) {
	t.Parallel()
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: ``,
			args: args{
				grid: [][]int{
					{2, 0, 0, 1},
					{0, 3, 1, 0},
					{0, 5, 2, 0},
					{4, 0, 0, 2},
				},
			},
			want: true,
		},
		{
			name: ``,
			args: args{
				grid: [][]int{
					{5, 7, 0},
					{0, 3, 1},
					{0, 5, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkXMatrix(tt.args.grid); got != tt.want {
				t.Errorf("checkXMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
