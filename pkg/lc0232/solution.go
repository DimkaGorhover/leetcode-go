package lc0232

type MyQueue struct {
	s1, s2 stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: newStack(),
		s2: newStack(),
	}
}

func (q *MyQueue) Push(x int) {
	q.s1.push(x)
}

func (q *MyQueue) Pop() int {
	q.rebalance()
	value, _ := q.s2.pop()
	return value
}

func (q *MyQueue) Peek() int {
	q.rebalance()
	value, _ := q.s2.peek()
	return value
}

func (q *MyQueue) rebalance() {
	if !q.s2.empty() {
		return
	}
	for {
		v, found := q.s1.pop()
		if !found {
			break
		}
		q.s2.push(v)
	}
}

func (q *MyQueue) Empty() bool {
	return q.s1.empty() && q.s2.empty()
}

type stack interface {
	push(value int)
	pop() (int, bool)
	peek() (int, bool)
	empty() bool
}

type stackImpl struct {
	arr    []int
	cursor int
}

func newStack() stack {
	return &stackImpl{}
}

func (s *stackImpl) empty() bool {
	return s.cursor == 0
}

func (s *stackImpl) push(value int) {
	if s.cursor < len(s.arr) {
		s.arr[s.cursor] = value
	} else {
		s.arr = append(s.arr, value)
	}
	s.cursor++
}

func (s *stackImpl) peek() (int, bool) {
	if s.cursor == 0 {
		return 0, false
	}
	return s.arr[s.cursor-1], true
}

func (s *stackImpl) pop() (int, bool) {
	if s.cursor == 0 {
		return 0, false
	}
	s.cursor--
	value := s.arr[s.cursor]
	return value, true
}
