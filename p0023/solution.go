package main

import (
	"container/heap"
	"fmt"
	. "github.com/DimkaGorhover/leetcode-go/common"
)

type priorityQueue struct {
	nodes []*ListNode
}

func (pq *priorityQueue) Len() int {
	return len(pq.nodes)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (pq.nodes)[i].Val < (pq.nodes)[j].Val
}

func (pq *priorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	pq.nodes = append(pq.nodes, node)
}

func (pq *priorityQueue) Pop() interface{} {
	l := pq.Len()
	if l == 0 {
		return nil
	}
	el := pq.nodes[l-1]
	pq.nodes = pq.nodes[:l-1]
	return el
}

func (pq *priorityQueue) Swap(i, j int) {
	tmp := pq.nodes[i]
	pq.nodes[i] = pq.nodes[j]
	pq.nodes[j] = tmp
}

func (pq *priorityQueue) String() string {
	str := ""
	prefix := ""
	for _, node := range pq.nodes {
		str = fmt.Sprintf("%s%s%v", str, prefix, node)
		prefix = ", "
	}
	return "[" + str + "]"
}

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	h := &priorityQueue{}
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	tmp := &ListNode{}
	head := tmp
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		head.Next = node
		head = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return tmp.Next
}
