package p0003

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `aa`,
			args: args{
				s: `aa`,
			},
			want: 1,
		},
		{
			name: `abcabcbb`,
			args: args{
				s: `abcabcbb`,
			},
			want: 3,
		},
		{
			name: `bbbbb`,
			args: args{
				s: `bbbbb`,
			},
			want: 1,
		},
		{
			name: `pwwkew`,
			args: args{
				s: `pwwkew`,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
