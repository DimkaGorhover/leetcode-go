package lc0098

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	t.Parallel()
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `test_001`,
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: `test_002`,
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: false,
		},
		{
			name: `test_003`,
			args: args{
				root: &TreeNode{Val: 5},
			},
			want: true,
		},
		{
			name: `test_004`,
			args: args{
				root: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidBST(tt.args.root))
		})
	}
}
