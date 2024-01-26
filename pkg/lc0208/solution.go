package lc0208

const (
	offset   = 'a'
	capacity = 'z' - offset + 1
)

type Trie struct {
	root  *node
	words map[string]bool
}

func Constructor() Trie {
	return Trie{
		root:  &node{},
		words: map[string]bool{},
	}
}

func (t *Trie) Insert(word string) {
	t.words[word] = true
	if len(word) == 0 {
		return
	}
	n := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - offset
		if n.children[index] == nil {
			n.children[index] = &node{}
		}
		n = n.children[index]
	}
	n.word = true
}

func (t *Trie) Search(word string) bool {
	return t.words[word]
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.root.findNode(prefix) != nil
}

type node struct {
	children [capacity]*node
	word     bool
}

func (n *node) findNode(prefix string) *node {
	current := n
	for i := 0; i < len(prefix); i++ {
		if current == nil {
			return nil
		}
		current = current.children[prefix[i]-offset]
	}
	return current
}
