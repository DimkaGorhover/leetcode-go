package main

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	t.Parallel()
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				lists: []*ListNode{
					ListNodeOf(1, 2, 3, 9),
					ListNodeOf(1, 2, 3, 8),
					ListNodeOf(1, 2, 3, 4),
					ListNodeOf(1, 2, 3, 7),
					ListNodeOf(1, 2, 3, 5),
				},
			},
			want: ListNodeOf(1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 5, 7, 8, 9),
		},
		{
			name: "",
			args: args{
				lists: []*ListNode{
					ListNodeOf(),
					ListNodeOf(),
				},
			},
			want: ListNodeOf(),
		},
		{
			name: "",
			args: args{},
			want: ListNodeOf(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
