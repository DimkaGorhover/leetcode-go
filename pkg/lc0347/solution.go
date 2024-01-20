package lc0347

func topKFrequent(nums []int, k int) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	head := newCountNode(-1)
	tail := newCountNode(-1)
	connectCountNodes(head, tail)
	head.insertNext(newCountNode(0))
	nodeMap := make(map[int]*node, len(nums))

	for _, num := range nums {

		if n, found := nodeMap[num]; !found {
			n = newNode(num, nil)
			head.getNextWithCount(0).add(n)
			nodeMap[num] = n
		} else {
			parent := n.parent
			n.unlink()
			parent.getNext().add(n)
			if parent.isEmpty() {
				parent.unlink()
			}
		}
	}

	result := make([]int, k)
	itr := newDownIterator(tail)
	for i := 0; i < k; i++ {
		if n, found := itr.next(); found {
			result[i] = n.value
		} else {
			break
		}
	}

	return result
}

type countNode struct {
	count      int
	prev, next *countNode
	head, tail *node
}

func newCountNode(count int) *countNode {
	c := &countNode{
		count: count,
	}
	c.head = newNode(-1, c)
	c.tail = newNode(-1, c)
	connectNodes(c.head, c.tail)
	return c
}

func connectCountNodes(prev, next *countNode) {
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}

func (cn *countNode) isEmpty() bool {
	return cn.head.next == cn.tail
}

func (cn *countNode) getNext() *countNode {
	return cn.getNextWithCount(cn.count + 1)
}

func (cn *countNode) getNextWithCount(count int) *countNode {
	if cn.next.count == count {
		return cn.next
	}
	newCN := newCountNode(count)
	cn.insertNext(newCN)
	return newCN
}

func (cn *countNode) insertNext(node *countNode) {
	next := cn.next
	connectCountNodes(cn, node)
	connectCountNodes(node, next)
}

func (cn *countNode) add(newNode *node) {
	newNode.parent = cn
	cn.head.connectNext(newNode)
}

func (cn *countNode) unlink() {
	connectCountNodes(cn.prev, cn.next)
	cn.prev, cn.next = nil, nil
}

type node struct {
	value      int
	parent     *countNode
	prev, next *node
}

func newNode(value int, parent *countNode) *node {
	return &node{
		value:  value,
		parent: parent,
	}
}

func connectNodes(prev, next *node) {
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}

func (n *node) connectNext(newNode *node) {
	newNode.parent = n.parent
	next := n.next
	connectNodes(n, newNode)
	connectNodes(newNode, next)
}

func (n *node) unlink() {
	connectNodes(n.prev, n.next)
	n.prev, n.next, n.parent = nil, nil, nil
}

type downIterator struct {
	cn *countNode
	n  *node
}

func newDownIterator(cn *countNode) *downIterator {
	return &downIterator{cn: cn}
}

func (i *downIterator) hasNext() bool {
	if i.cn == nil {
		return false
	}
	if i.n == nil {
		i.n = i.cn.tail.prev
	}
	if i.n != i.cn.head {
		return true
	}
	i.n = nil
	i.cn = i.cn.prev
	return i.hasNext()
}

func (i *downIterator) next() (*node, bool) {
	if !i.hasNext() {
		return nil, false
	}
	v := i.n
	i.n = i.n.prev
	return v, true
}
