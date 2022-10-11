package p0199

import (
	. "github.com/DimkaGorhover/leetcode-go/common"
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	t.Parallel()
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: ``,
			args: args{
				root: NewTreeNodeBuilder().AddAll(1, 2, 3).AddNil().Add(5).AddNil().Add(4).Build(),
			},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
