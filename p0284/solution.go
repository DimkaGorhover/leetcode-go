package p0284

type PeekingIterator struct {
	iter      *Iterator
	nextValue int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:      iter,
		nextValue: 0,
	}
}

func (i *PeekingIterator) hasNext() bool {
	if i.nextValue > 0 {
		return true
	}
	if !i.iter.hasNext() {
		return false
	}
	i.nextValue = i.iter.next()
	return true
}

func (i *PeekingIterator) next() int {
	value := i.peek()
	i.nextValue = 0
	return value
}

func (i *PeekingIterator) peek() int {
	if !i.hasNext() {
		panic("no more elements")
	}
	return i.nextValue
}

// ============================================================================
// Test Data
//

type Iterator struct {
	arr    []int
	cursor int
}

func NewIterator(arr []int) *Iterator {
	return &Iterator{
		arr:    arr,
		cursor: 0,
	}
}

func (i *Iterator) hasNext() bool {
	return i.cursor < len(i.arr)
}

func (i *Iterator) next() int {
	if !i.hasNext() {
		panic("no more elements")
	}
	value := i.arr[i.cursor]
	i.cursor++
	return value
}
