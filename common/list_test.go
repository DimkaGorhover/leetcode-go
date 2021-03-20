package common

import (
	"testing"
)

func TestListNode_Length(t *testing.T) {
	tests := []struct {
		name string
		list *ListNode
		want uint32
	}{
		{name: "", list: ListNodeOf(1, 2, 3, 4, 5), want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	tests := []struct {
		name string
		list *ListNode
		want string
	}{
		{name: "", list: ListNodeOf(1, 2, 3, 4, 5), want: "[1, 2, 3, 4, 5]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
