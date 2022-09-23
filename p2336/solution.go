package p2336

import "container/heap"

type SmallestInfiniteSet struct {
	set  []bool
	heap *priorityQueue
	next int
}

func Constructor() SmallestInfiniteSet {
	setCapacity := 1_001
	set := make([]bool, setCapacity, setCapacity)
	return SmallestInfiniteSet{
		set:  set,
		heap: newPriorityQueue(),
		next: 1,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.heap.Len() > 0 {
		num := heap.Pop(s.heap).(int)
		s.set[num] = false
		return num
	}
	next := s.next
	s.next++
	return next
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if s.set[num] || num >= s.next {
		return
	}
	s.set[num] = true
	heap.Push(s.heap, num)
}

// ============================================================================

type priorityQueue struct {
	heap.Interface
	numbers []int
	cursor  int
}

func newPriorityQueue() *priorityQueue {
	numbers := make([]int, 1<<10, 1<<10)
	h := priorityQueue{
		numbers: numbers,
		cursor:  0,
	}
	heap.Init(&h)
	return &h
}

func (pq *priorityQueue) Len() int {
	return pq.cursor
}

func (pq *priorityQueue) grow() {
	currentCapacity := cap(pq.numbers)
	if currentCapacity <= pq.cursor {
		newCapacity := currentCapacity << 1
		newArr := make([]int, newCapacity, newCapacity)
		copy(pq.numbers, newArr)
		pq.numbers = newArr
	}
}

func (pq *priorityQueue) Less(i, j int) bool {
	return pq.numbers[i] < pq.numbers[j]
}

func (pq *priorityQueue) Push(x interface{}) {
	pq.grow()
	pq.numbers[pq.cursor] = x.(int)
	pq.cursor++
}

func (pq *priorityQueue) Pop() interface{} {
	if pq.cursor == 0 {
		return -1
	}
	pq.cursor--
	num := pq.numbers[pq.cursor]
	return num
}

func (pq *priorityQueue) Swap(i, j int) {
	tmp := pq.numbers[i]
	pq.numbers[i] = pq.numbers[j]
	pq.numbers[j] = tmp
}
