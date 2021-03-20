package p0203

import . "github.com/DimkaGorhover/leetcode-go/common"

func removeElements(head *ListNode, val int) *ListNode {
	tmpHead := &ListNode{Next: head}
	head = tmpHead
	for head != nil {
		next := head.Next
		if next != nil && next.Val == val {
			head.Next = next.Next
		} else {
			head = next
		}
	}
	return tmpHead.Next
}
