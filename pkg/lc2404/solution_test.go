package lc2404

import "testing"

func Test_mostFrequentEven(t *testing.T) {
	t.Parallel()
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: ``,
			args: args{
				nums: []int{0, 1, 2, 2, 4, 4, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequentEven(tt.args.nums); got != tt.want {
				t.Errorf("mostFrequentEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
