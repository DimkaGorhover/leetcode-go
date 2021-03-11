package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LLHeap []*ListNode

func (L *LLHeap) Len() int {
	return len(*L)
}

func (L *LLHeap) Less(i, j int) bool {
	return (*L)[i].Val < (*L)[j].Val
}

func (L *LLHeap) Push(x interface{}) {
	*L = append(*L, x.(*ListNode))
}

func (L *LLHeap) Pop() interface{} {
	old := *L
	l := len(old)
	last := old[l-1]
	*L = old[:l-1]
	return last
}

func (L LLHeap) Swap(i, j int) {
	tmp := L[i]
	L[i] = L[j]
	L[j] = tmp
}

func mergeKLists(lists []*ListNode) *ListNode {

	h := make(LLHeap, 0)
	heap.Init(&h)

	for i := 0; i < len(lists); i = i + 1 {
		head := *(lists[i])
		heap.Push(&h, &head)
	}

	for len(h) > 0 {
		fmt.Printf(" --> %d\n", heap.Pop(&h))
	}

	return nil
}

func main() {

	nodes := []*ListNode{
		&(ListNode{Val: 4}),
		&(ListNode{Val: 2}),
		&(ListNode{Val: 1}),
		&(ListNode{Val: 14}),
	}

	mergeKLists(nodes)
}
