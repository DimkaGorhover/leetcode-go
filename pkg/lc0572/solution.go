package lc0572

import . "github.com/DimkaGorhover/leetcode-go/pkg/common"

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot == nil {
		return false
	}
	return dfs(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func dfs(root, subRoot *TreeNode) bool {
	if subRoot == nil {
		return root == nil
	}
	if root == nil {
		return false
	}
	return root.Val == subRoot.Val && dfs(root.Left, subRoot.Left) && dfs(root.Right, subRoot.Right)
}
