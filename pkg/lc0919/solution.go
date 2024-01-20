package lc0919

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
)

type CBTInserter struct {
	root            *TreeNode
	prev, current   []*TreeNode
	prevI, currentI int
}

// TODO: is not solved yet
func Constructor(root *TreeNode) CBTInserter {
	if root == nil {
		panic(`tree node is nil`)
	}
	prev := make([]*TreeNode, 1)
	current := make([]*TreeNode, 2)
	cbt := CBTInserter{
		root:     &TreeNode{Val: root.Val},
		prev:     prev,
		prevI:    0,
		current:  current,
		currentI: 0,
	}
	cbt.populate()
	return cbt
}

func (cbt *CBTInserter) Get_root() *TreeNode {
	return cbt.root
}

func (cbt *CBTInserter) populate() {
	cbt.prev[0] = cbt.root

	q := newQueue[*TreeNode]()
	if cbt.root.Left != nil {
		q.Push(cbt.root.Left)
	}
	if cbt.root.Right != nil {
		q.Push(cbt.root.Right)
	}
	for q.Size() > 0 {
		s := q.Size()
		for i := 0; i < s; i++ {
			node := q.Poll()
			cbt.Insert(node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
}

func (cbt *CBTInserter) Insert(val int) int {
	if cbt.currentI == len(cbt.current) {
		cbt.prev = cbt.current
		newCapacity := len(cbt.current) << 1
		newArr := make([]*TreeNode, newCapacity, newCapacity)
		cbt.current = newArr
		cbt.prevI = 0
		cbt.currentI = 0
	}

	parent := cbt.prev[cbt.prevI]
	if parent.Left == nil {
		parent.Left = &TreeNode{Val: val}
		cbt.current[cbt.currentI] = parent.Left
		cbt.currentI++
		return parent.Val
	}

	parent.Right = &TreeNode{Val: val}
	cbt.current[cbt.currentI] = parent.Right
	cbt.currentI++
	return parent.Val
}
