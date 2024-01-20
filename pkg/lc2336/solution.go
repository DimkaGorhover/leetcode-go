package lc2336

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

func (set *SmallestInfiniteSet) PopSmallest() int {
	if set.heap.Len() > 0 {
		num := heap.Pop(set.heap).(int)
		set.set[num] = false
		return num
	}
	next := set.next
	set.next++
	return next
}

func (set *SmallestInfiniteSet) AddBack(num int) {
	if set.set[num] || num >= set.next {
		return
	}
	set.set[num] = true
	heap.Push(set.heap, num)
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
