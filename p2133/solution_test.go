package p2133

import "testing"

func Test_checkValid(t *testing.T) {
	t.Parallel()
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: ``,
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{3, 1, 2},
					{2, 3, 1},
				},
			},
			want: true,
		},
		{
			name: ``,
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValid(tt.args.matrix); got != tt.want {
				t.Errorf("checkValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
