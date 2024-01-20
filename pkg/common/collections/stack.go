package collections

import (
	"fmt"
)

type Collection[T any] interface {
	HasElements() bool
	IsEmpty() bool
}

type Stack[T any] interface {
	Collection[T]
	Push(value T)
	Pop() T
}

type stack[T any] struct {
	content []T
	cursor  int
}

func NewStack[T any](capacity int) Stack[T] {
	return &stack[T]{
		content: make([]T, capacity),
		cursor:  0,
	}
}

func (s *stack[T]) Push(value T) {
	capacity := cap(s.content)
	if capacity <= s.cursor {
		newCapacity := capacity << 1
		newContent := make([]T, newCapacity, newCapacity)
		copy(newContent, s.content)
		s.content = newContent
	}
	s.content[s.cursor] = value
	s.cursor++
}

func (s *stack[T]) Pop() T {
	if s.cursor == 0 {
		panic(`stack is empty`)
	}
	s.cursor--
	return s.content[s.cursor]
}

func (s *stack[T]) String() string {
	return fmt.Sprintf("%v", s.content)
}

func (s *stack[T]) HasElements() bool {
	return s.cursor > 0
}

func (s *stack[T]) IsEmpty() bool {
	return s.cursor == 0
}
