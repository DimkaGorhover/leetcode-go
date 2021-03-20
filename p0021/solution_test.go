package p0021

import (
	. "github.com/DimkaGorhover/leetcode-go/common"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "", args: args{
			l1: ListNodeOf(1, 3, 5),
			l2: ListNodeOf(2, 4, 6),
		}, want: ListNodeOf(1, 2, 3, 4, 5, 6)},
		{name: "", args: args{
			l1: nil,
			l2: ListNodeOf(2, 4, 6),
		}, want: ListNodeOf(2, 4, 6)},
		{name: "", args: args{
			l1: ListNodeOf(1, 3, 5),
			l2: nil,
		}, want: ListNodeOf(1, 3, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
