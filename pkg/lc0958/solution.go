package lc0958

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
)

//goland:noinspection GoUnusedFunction
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := newQueue()
	q.Add(root)
	foundNull := false

	for !q.IsEmpty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			n := q.Poll()
			if n == nil {
				foundNull = true
				continue
			}
			if foundNull {
				return false
			}
			q.Add(n.Left)
			q.Add(n.Right)
		}
	}

	return true
}

type queue interface {
	IsEmpty() bool
	Size() int
	Poll() *TreeNode
	Add(value *TreeNode)
}

type twoPointersArrayQueue struct {
	arr  []*TreeNode
	left int
}

func newQueue() queue {
	if true {
		return newTwoPointersQueue()
	}
	return newDoublyLinkedList()
}

func newTwoPointersQueue() queue {
	return &twoPointersArrayQueue{}
}

func (t *twoPointersArrayQueue) Size() int {
	return len(t.arr) - t.left
}

func (t *twoPointersArrayQueue) IsEmpty() bool {
	return t.Size() == 0
}

func (t *twoPointersArrayQueue) Poll() *TreeNode {
	if t.IsEmpty() {
		return nil
	}
	value := t.arr[t.left]
	t.left++
	return value
}

func (t *twoPointersArrayQueue) Add(value *TreeNode) {
	t.arr = append(t.arr, value)
}

type doublyLinkedQueue struct {
	head, tail *doublyLinkedListNode
	size       int
}

func newDoublyLinkedList() queue {
	head := newNode(nil)
	tail := newNode(nil)
	head.next = tail
	tail.prev = head
	return &doublyLinkedQueue{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (q *doublyLinkedQueue) Size() int {
	return q.size
}

func (q *doublyLinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *doublyLinkedQueue) Add(value *TreeNode) {
	q.tail.prev.insertNext(value)
	q.size++
}

func (q *doublyLinkedQueue) Poll() *TreeNode {
	if q.size > 0 {
		q.size--
		n := q.head.next
		n.unlink()
		return n.value
	}
	return nil
}

type doublyLinkedListNode struct {
	value      *TreeNode
	prev, next *doublyLinkedListNode
}

func (n *doublyLinkedListNode) unlink() {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func newNode(value *TreeNode) *doublyLinkedListNode {
	return &doublyLinkedListNode{
		value: value,
	}
}

func (n *doublyLinkedListNode) insertNext(value *TreeNode) *doublyLinkedListNode {
	next := n.next
	n.next = newNode(value)
	n.next.next = next
	n.next.prev = n
	if next.prev != nil {
		next.prev = n.next
	}
	return n.next
}
