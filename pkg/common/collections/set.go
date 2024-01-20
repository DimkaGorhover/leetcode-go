package collections

import (
	"fmt"
	"strings"
	"sync"
)

type Set[T comparable] interface {
	fmt.Stringer
	Size() int
	Contains(T) bool
	Add(T) bool
	Remove(T) bool
	AddAll(...T) bool
	Values() []T
}

type set[T comparable] struct {
	head, tail *DoublyLinkedNode[T]
	table      map[T]*DoublyLinkedNode[T]
	size       int
	lock       *sync.Mutex
}

func NewSet[T comparable]() Set[T] {
	var lock sync.Mutex
	head := &DoublyLinkedNode[T]{}
	head.InsertNext(&DoublyLinkedNode[T]{})
	return &set[T]{
		head:  head,
		tail:  head.Next,
		table: map[T]*DoublyLinkedNode[T]{},
		lock:  &lock,
		size:  0,
	}
}

func (u *set[T]) Size() int {
	return u.size
}

func (u *set[T]) Contains(value T) bool {
	node, exists := u.table[value]
	return exists && node != nil
}

func (u *set[T]) Remove(value T) bool {
	u.lock.Lock()
	defer u.lock.Unlock()
	if u.size == 0 {
		return false
	}
	node, exists := u.table[value]
	if !exists || node == nil {
		return false
	}
	node.Unlink()
	delete(u.table, value)
	u.size--
	return true
}

func (u *set[T]) Add(value T) bool {
	u.lock.Lock()
	defer u.lock.Unlock()
	return u.add(value)
}

func (u *set[T]) add(value T) bool {
	if v, exists := u.table[value]; exists && v != nil {
		return false
	}

	node := NewDoublyLinkedNode[T](value)
	u.tail.InsertPrev(node)

	u.table[value] = node
	u.size++

	return true
}

func (u *set[T]) AddAll(values ...T) bool {
	u.lock.Lock()
	defer u.lock.Unlock()
	result := false
	for _, value := range values {
		result = u.add(value) || result
	}
	return result
}

func (u *set[T]) Values() []T {
	u.lock.Lock()
	defer u.lock.Unlock()
	newSlice := make([]T, u.size)
	node := u.head.Next
	for i := 0; i < u.size; i++ {
		newSlice[i] = node.Value
		node = node.Next
	}
	return newSlice
}

func (u *set[T]) String() string {
	u.lock.Lock()
	defer u.lock.Unlock()
	if u.size == 0 {
		return `[]`
	}
	var sb strings.Builder
	sb.WriteString(`[`)
	node := u.head.Next
	sep := ``
	for i := 0; node != nil && i < u.size; i++ {
		sb.WriteString(sep)
		sb.WriteString(node.String())
		node = node.Next
		sep = `, `
	}
	sb.WriteString(`]`)
	return sb.String()
}
