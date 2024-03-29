package common

import (
	"fmt"
)

// ListNode - ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeOf - ListNodeOf
func ListNodeOf(numbers ...int) *ListNode {
	tmp := &ListNode{}
	head := tmp
	for _, number := range numbers {
		head.Next = &ListNode{Val: number}
		head = head.Next
	}
	return tmp.Next
}

// Length - Length
func (node *ListNode) Length() uint32 {
	var length uint32 = 1
	next := node.Next
	for next != nil {
		length++
		next = next.Next
	}
	return length
}

// String - String
func (node *ListNode) String() string {
	sb := fmt.Sprintf("%d", node.Val)
	next := node.Next
	for next != nil {
		sb = fmt.Sprintf("%s, %d", sb, next.Val)
		next = next.Next
	}
	return fmt.Sprintf("[%s]", sb)
}

func (node *ListNode) Equals(other *ListNode) bool {
	if node == other {
		return true
	}
	n1 := node
	n2 := other
	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1 == n2
}
