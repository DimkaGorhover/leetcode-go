package p1137

import "testing"

func Test_tribonacci(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Name n=0 -> 0, ", args: args{n: 0}, want: 0},
		{name: "Name n=1 -> 1, ", args: args{n: 1}, want: 1},
		{name: "Name n=2 -> 1, ", args: args{n: 2}, want: 1},
		{name: "Name n=3 -> 2, ", args: args{n: 3}, want: 2},
		{name: "Name n=4 -> 4, ", args: args{n: 4}, want: 4},
		{name: "Name n=5 -> 7, ", args: args{n: 5}, want: 7},
		{name: "Name n=6 -> 13, ", args: args{n: 6}, want: 13},
		{name: "Name n=7 -> 24, ", args: args{n: 7}, want: 24},
		{name: "Name n=8 -> 44, ", args: args{n: 8}, want: 44},
		{name: "Name n=9 -> 81, ", args: args{n: 9}, want: 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
