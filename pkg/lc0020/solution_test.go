package lc0021

import "testing"

func Test_isValid(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}

	longTestCase0 := "[([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([([()])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])])]"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "()", args: args{"()"}, want: true},
		{name: "()[]{}", args: args{"()[]{}"}, want: true},
		{name: "{[]}", args: args{"{[]}"}, want: true},
		{name: "(]", args: args{"(]"}, want: false},
		{name: "Long test Case 0", args: args{longTestCase0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}