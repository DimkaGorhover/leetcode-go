package p0203

import (
	. "github.com/DimkaGorhover/leetcode-go/common"
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	t.Parallel()
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: ListNodeOf(1, 2, 6, 3, 4, 5, 6),
				val:  6,
			},
			want: ListNodeOf(1, 2, 3, 4, 5),
		},
		{
			name: "",
			args: args{
				head: ListNodeOf(7, 7, 7, 7),
				val:  7,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: ListNodeOf(1, 2, 2, 1),
				val:  2,
			},
			want: ListNodeOf(1, 1),
		},
		{
			name: "",
			args: args{
				head: ListNodeOf(1, 2, 2, 1),
				val:  1,
			},
			want: ListNodeOf(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
