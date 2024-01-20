package collections

import "fmt"

type DoublyLinkedNode[T interface{}] struct {
	Value      T
	Prev, Next *DoublyLinkedNode[T]
}

func (n *DoublyLinkedNode[T]) String() string {
	return fmt.Sprintf(`%v`, n.Value)
}

func NewDoublyLinkedNode[T interface{}](value T) *DoublyLinkedNode[T] {
	return &DoublyLinkedNode[T]{Value: value}
}

func (n *DoublyLinkedNode[T]) Unlink() {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
}

func (n *DoublyLinkedNode[T]) InsertPrev(node *DoublyLinkedNode[T]) {
	prev := n.Prev
	node.Next = n
	node.Prev = prev
	n.Prev = node
	if prev != nil {
		prev.Next = node
	}
}

func (n *DoublyLinkedNode[T]) InsertNext(node *DoublyLinkedNode[T]) {
	next := n.Next
	node.Prev = n
	node.Next = next
	n.Next = node
	if next != nil {
		next.Prev = node
	}
}
