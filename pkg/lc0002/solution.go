package lc0001

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
)

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	rest := 0
	for l1 != nil || l2 != nil {
		sum := val(l1) + val(l2) + rest
		rest = sum / 10
		tail = swap(tail, sum%10)
		l1 = next(l1)
		l2 = next(l2)
	}
	if rest > 0 {
		tail = swap(tail, rest)
	}
	return head.Next
}

func swap(node *ListNode, value int) *ListNode {
	next := &ListNode{Val: value}
	node.Next = next
	return next
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}
	return nil
}

func val(node *ListNode) int {
	if node != nil {
		return node.Val
	}
	return 0
}
