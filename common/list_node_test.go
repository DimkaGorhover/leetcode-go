package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNode_Length(t *testing.T) {
	t.Parallel()
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
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	t.Parallel()
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
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestListNode_Equals(t *testing.T) {
	t.Parallel()
	type args struct {
		list  *ListNode
		other *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				list:  ListNodeOf(1, 2, 3, 4, 5),
				other: ListNodeOf(1, 2, 3, 4, 5),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				list:  ListNodeOf(1, 2, 3, 4, 5, 5),
				other: ListNodeOf(1, 2, 3, 4, 5),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				list:  ListNodeOf(1, 2, 3, 4, 1),
				other: ListNodeOf(1, 2, 3, 4, 5),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				list:  nil,
				other: ListNodeOf(1, 2, 3, 4, 5),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				list:  ListNodeOf(1, 2, 3, 4, 5),
				other: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.list.Equals(tt.args.other); got != tt.want {
				if tt.want {
					t.Errorf(`lists must be equal, but they don't`)
				} else {
					t.Errorf(`lists must not be equal, but they do`)
				}
			}
		})
	}
}
