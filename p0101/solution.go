package p0101

import . "github.com/DimkaGorhover/leetcode-go/common"

func isSymmetric(root *TreeNode) bool {
	return isSame(root, root)
}

func isSame(left, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	return right != nil &&
		left.Val == right.Val &&
		isSame(left.Left, right.Right) &&
		isSame(left.Right, right.Left)
}
