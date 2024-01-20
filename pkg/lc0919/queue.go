package lc0919

type Queue[T any] interface {
	Size() int
	Push(T)
	PushAll(...T)
	Poll() T
	Peek() T
	Values() []T
}

func newQueue[T any]() Queue[T] {
	return &queue[T]{}
}

type queue[T any] struct {
	head, tail *queueNode[T]
	size       int
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *queue[T]) Push(value T) {
	if q.size == 0 {
		q.head = newQueueNode(value)
		q.tail = q.head
		q.size++
		return
	}
	q.tail.next = newQueueNode(value)
	q.tail = q.tail.next
	q.size++
}

func (q *queue[T]) Poll() T {
	if q.size == 0 {
		panic(`queue is empty`)
	}
	val := q.head.val
	q.head = q.head.next
	q.size--
	return val
}

func (q *queue[T]) Peek() T {
	if q.size == 0 {
		panic(`queue is empty`)
	}
	return q.head.val
}

func (q *queue[T]) Values() []T {
	arr := make([]T, q.size)
	if q.size > 0 {
		node := q.head
		for i := 0; i < q.size; i++ {
			arr[i] = node.val
			node = node.next
		}
	}
	return arr
}

type queueNode[T any] struct {
	val  T
	next *queueNode[T]
}

func newQueueNode[T any](val T) *queueNode[T] {
	return &queueNode[T]{val: val}
}
