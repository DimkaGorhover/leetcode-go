package lc0098

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
	"math"
)

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if val >= max || val <= min {
		return false
	}
	return helper(root.Left, min, val) && helper(root.Right, val, max)
}
