package lc0199

import . "github.com/DimkaGorhover/leetcode-go/pkg/common"

// rightSideView - 199. Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/submissions/
func rightSideView(root *TreeNode) []int {
	return traverse(root, []int{}, 0)
}

func traverse(node *TreeNode, list []int, level int) []int {
	if node == nil {
		return list
	}
	if level < len(list) {
		list[level] = node.Val
	} else {
		list = append(list, node.Val)
	}
	list = traverse(node.Left, list, level+1)
	list = traverse(node.Right, list, level+1)
	return list
}
