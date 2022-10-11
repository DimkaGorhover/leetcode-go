package p0001

import (
	. "github.com/DimkaGorhover/leetcode-go/common"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Parallel()
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "asd",
			args: args{
				l1: ListNodeOf(2, 4, 3),
				l2: ListNodeOf(5, 6, 4),
			},
			want: ListNodeOf(7, 0, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !got.Equals(tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
