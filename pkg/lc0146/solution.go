package lc0146

type LRUCache struct {
	c CacheProvider
}

func Constructor(capacity int) LRUCache {
	var cp CacheProvider
	switch capacity {
	case 0:
		cp = EmptyCacheProvider(true)
	case 1:
		cp = NewSingleValueCacheProvider()
	default:
		cp = NewDefaultCacheProvider(capacity)
	}

	return LRUCache{cp}
}

func (c *LRUCache) Get(key int) int {
	return c.c.Get(key)
}

func (c *LRUCache) Put(key int, value int) {
	c.c.Put(key, value)
}

type CacheProvider interface {
	Get(int) int
	Put(int, int)
}

type EmptyCacheProvider bool

func (c EmptyCacheProvider) Get(_ int) int {
	return -1
}

func (c EmptyCacheProvider) Put(_ int, _ int) {}

type singleValueCacheProvider struct {
	init       bool
	key, value int
}

func NewSingleValueCacheProvider() CacheProvider {
	return &singleValueCacheProvider{}
}

func (c *singleValueCacheProvider) Get(key int) int {
	if c.init && c.key == key {
		return c.value
	}
	return -1
}

func (c *singleValueCacheProvider) Put(key int, value int) {
	c.key = key
	c.value = value
	c.init = true
}

type defaultCacheProvider struct {
	head, tail       *Node
	cache            [10001]*Node
	capacity, length int
}

func NewDefaultCacheProvider(capacity int) CacheProvider {
	head := NewNode(0, 0)
	tail := NewNode(0, 0)
	head.insertNext(tail)
	return &defaultCacheProvider{
		head:     head,
		tail:     tail,
		capacity: capacity,
		length:   0,
		cache:    [10001]*Node{},
	}
}

func (c *defaultCacheProvider) Get(key int) int {
	if node := c.cache[key]; node != nil {
		node.unlink()
		c.head.insertNext(node)
		return node.value
	}
	return -1
}

func (c *defaultCacheProvider) Put(key int, value int) {
	if node := c.cache[key]; node != nil {
		node.unlink()
		c.head.insertNext(node)
		node.value = value
		return
	}

	if c.length == c.capacity {
		node := c.tail.prev
		node.unlink()
		c.cache[node.key] = nil
		c.length--
	}

	node := NewNode(key, value)
	c.head.insertNext(node)
	c.cache[key] = node

	c.length++
}

type Node struct {
	key, value int
	prev, next *Node
}

func NewNode(key int, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (n *Node) unlink() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.prev = nil
	n.next = nil
}

func (n *Node) insertNext(node *Node) {
	if n == node {
		panic(`cannot link itself`)
	}
	if node == nil {
		panic(`do not insert nil nodes`)
	}
	next := n.next
	n.next = node
	node.prev = n
	node.next = next
	if next != nil {
		next.prev = node
	}
}
