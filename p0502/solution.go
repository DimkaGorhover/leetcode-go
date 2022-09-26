package p0502

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k, w int, profits []int, capital []int) int {
	if k <= 0 {
		return w
	}
	if isFastSolution(w, capital) {
		return fastSolution(k, w, profits)
	}
	return twoHeapsSolution(k, w, profits, capital)
}

func isFastSolution(w int, capital []int) bool {
	for _, c := range capital {
		if c > w {
			return false
		}
	}
	return true
}

func fastSolution(k, w int, profits []int) int {
	sort.Ints(profits)
	for i := len(profits) - 1; i >= 0 && k > 0; i-- {
		w += profits[i]
		k--
	}
	return w
}

func twoHeapsSolution(k, w int, profits []int, capital []int) int {

	pHeap := newProjectsHeap(projectsHeapProps{
		capacity:   len(profits),
		comparator: lessCapitalComparator,
	})

	for i := 0; i < len(profits); i++ {
		heap.Push(pHeap, project{
			capital: capital[i],
			profit:  profits[i],
		})
	}

	cHeap := newProjectsHeap(projectsHeapProps{
		capacity:   len(profits),
		comparator: maxProfitComparator,
	})

	for k > 0 {
		var p project
		for pHeap.Len() > 0 {
			p = heap.Pop(pHeap).(project)
			if p.capital <= w {
				heap.Push(cHeap, p)
			} else {
				heap.Push(pHeap, p)
				break
			}
		}

		if cHeap.Len() == 0 {
			break
		}

		w += heap.Pop(cHeap).(project).profit
		k--
	}

	return w
}

// ============================================================================

type project struct {
	capital, profit int
}

func lessCapitalComparator(left, right project) bool {
	if left.capital < right.capital {
		return true
	}
	if left.capital > right.capital {
		return false
	}
	if left.profit > right.profit {
		return true
	}
	return false
}

func maxProfitComparator(left, right project) bool {
	if left.profit > right.profit {
		return true
	}
	if left.profit < right.profit {
		return false
	}
	if left.capital < right.capital {
		return true
	}
	return false
}

type projectsHeap struct {
	arr        []project
	cursor     int
	comparator func(left, right project) bool
}

type projectsHeapProps struct {
	capacity   int
	comparator func(left, right project) bool
}

func newProjectsHeap(props projectsHeapProps) *projectsHeap {
	h := projectsHeap{
		arr:        make([]project, props.capacity, props.capacity),
		cursor:     0,
		comparator: props.comparator,
	}
	heap.Init(&h)
	return &h
}

func (pq *projectsHeap) Len() int {
	return pq.cursor
}

func (pq *projectsHeap) grow() {
	currentCapacity := cap(pq.arr)
	if currentCapacity <= pq.cursor {
		newCapacity := currentCapacity << 1
		newArr := make([]project, newCapacity, newCapacity)
		copy(pq.arr, newArr)
		pq.arr = newArr
	}
}

func (pq *projectsHeap) Less(i, j int) bool {
	return pq.comparator(pq.arr[i], pq.arr[j])
}

func (pq *projectsHeap) Push(x interface{}) {
	pq.grow()
	pq.arr[pq.cursor] = x.(project)
	pq.cursor++
}

func (pq *projectsHeap) Pop() interface{} {
	if pq.cursor == 0 {
		return -1
	}
	pq.cursor--
	num := pq.arr[pq.cursor]
	return num
}

func (pq *projectsHeap) Swap(i, j int) {
	tmp := pq.arr[i]
	pq.arr[i] = pq.arr[j]
	pq.arr[j] = tmp
}
