package lc0005

import "testing"

func Test_longestPalindrome(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `babad -> bab`,
			args: args{
				s: `babad`,
			},
			want: `bab`,
		},
		{
			name: `cbbd -> bb`,
			args: args{
				s: `cbbd`,
			},
			want: `bb`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
