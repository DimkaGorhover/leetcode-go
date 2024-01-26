package lc0046

import (
	"strconv"
	"strings"
)

func permute(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	if l == 1 {
		return [][]int{{nums[0]}}
	}
	if l == 2 {
		return [][]int{
			{nums[0], nums[1]},
			{nums[1], nums[0]},
		}
	}
	t := newTraverser(nums)
	t.traverse(0)
	return t.res
}

type traverser struct {
	res    permutations
	buffer intSlice
	ll     LinkedList
	cursor int
}

func newTraverser(nums []int) *traverser {
	capacity := len(nums)
	return &traverser{
		res:    make(permutations, fact(capacity)),
		buffer: newIntSlice(capacity),
		ll:     NewLinkedList(nums...),
		cursor: 0,
	}
}

func (t *traverser) traverse(level int) {

	if level == len(t.buffer) {
		t.res[t.cursor] = t.buffer.copy()
		t.cursor++
		return
	}

	for i := 0; i < len(t.buffer)-level; i++ {
		v, _ := t.ll.PollFirst()
		t.buffer[level] = v
		t.traverse(level + 1)
		t.ll.AddLast(v)
	}
}

func fact(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f = f * i
	}
	return f
}

type intSlice []int

func newIntSlice(capacity int) intSlice {
	return make(intSlice, capacity)
}

func (a intSlice) copy() []int {
	newArr := make([]int, len(a))
	copy(newArr, a)
	return newArr
}

type permutations [][]int

func (p permutations) String() string {
	var sb strings.Builder
	sep := [2]string{"", ""}
	sb.WriteByte('[')
	for _, a := range p {
		sb.WriteString(sep[0])
		sep[1] = ""
		sb.WriteByte('[')
		for _, b := range a {
			sb.WriteString(sep[1])
			sep[1] = ","
			sb.WriteString(strconv.Itoa(b))
		}
		sb.WriteByte(']')
		sep[0] = ","
	}
	sb.WriteByte(']')
	return sb.String()
}

type LinkedList interface {
	AddLast(int)
	AddFirst(int)
	PollFirst() (int, bool)
	PollLast() (int, bool)
}

func NewLinkedList(values ...int) LinkedList {
	head := &linkedListNode{}
	tail := &linkedListNode{}
	connect(head, tail)
	ll := &linkedList{
		head: head,
		tail: tail,
		size: 0,
	}
	for _, value := range values {
		ll.AddLast(value)
	}
	return ll
}

type linkedList struct {
	head, tail *linkedListNode
	size       int
}

func (l *linkedList) AddLast(value int) {
	node := &linkedListNode{value: value}
	head := l.tail.prev
	connect(head, node)
	connect(node, l.tail)
	l.size++
}

func (l *linkedList) AddFirst(value int) {
	node := &linkedListNode{value: value}
	tail := l.head.next
	connect(l.head, node)
	connect(node, tail)
	l.size++
}

func (l *linkedList) PollFirst() (int, bool) {
	if l.size == 0 {
		return 0, false
	}
	node := l.head.next
	tail := node.next
	connect(l.head, tail)
	l.size--
	return node.value, true
}

func (l *linkedList) PollLast() (int, bool) {
	if l.size == 0 {
		return 0, false
	}
	node := l.tail.prev
	prev := node.prev
	connect(l.tail, prev)
	l.size--
	return node.value, true
}

func (l *linkedList) String() string {
	if l.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteByte('[')
	sep := ""
	node := l.head.next
	for i := 0; i < l.size; i++ {
		sb.WriteString(sep)
		sep = ","
		sb.WriteString(node.String())
		node = node.next
	}
	sb.WriteByte(']')
	return sb.String()
}

type linkedListNode struct {
	value      int
	prev, next *linkedListNode
}

func (l linkedListNode) String() string {
	return strconv.Itoa(l.value)
}

func connect(prev, next *linkedListNode) {
	prev.next = next
	next.prev = prev
}
