package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Equals(other *TreeNode) bool {
	if n == other {
		return true
	}
	if other == nil {
		return false
	}
	return n.Val == other.Val && n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
}

type TreeNodeBuilder interface {
	Add(value int) TreeNodeBuilder
	AddAll(values ...int) TreeNodeBuilder
	AddNil() TreeNodeBuilder
	Build() *TreeNode
}

func NewTreeNodeBuilder() TreeNodeBuilder {
	return treeNodeBuilder{
		entries: []treeNodeBuilderEntry{},
	}
}

type treeNodeBuilderEntry struct {
	value int
	isNil bool
}

type treeNodeBuilder struct {
	entries []treeNodeBuilderEntry
}

func (t treeNodeBuilder) Add(value int) TreeNodeBuilder {
	t.entries = append(t.entries, treeNodeBuilderEntry{
		value: value,
		isNil: false,
	})
	return t
}

func (t treeNodeBuilder) AddAll(values ...int) TreeNodeBuilder {
	for _, value := range values {
		t.entries = append(t.entries, treeNodeBuilderEntry{
			value: value,
			isNil: false,
		})
	}
	return t
}

func (t treeNodeBuilder) AddNil() TreeNodeBuilder {
	t.entries = append(t.entries, treeNodeBuilderEntry{
		isNil: true,
	})
	return t
}

func (t treeNodeBuilder) Build() *TreeNode {

	if len(t.entries) == 0 {
		panic("Build: bad value")
	}

	if t.entries[0].isNil {
		panic("Build: bad value for the first element")
	}

	i := 1
	root := &TreeNode{Val: t.entries[0].value}
	prev := &[]*TreeNode{root}
	for len(*prev) > 0 && i < len(t.entries) {
		next := make([]*TreeNode, 0)
		for _, node := range *prev {
			if i < len(t.entries) && !t.entries[i].isNil {
				node.Left = &TreeNode{Val: t.entries[i].value}
				next = append(next, node.Left)
			}
			i++
			if i < len(t.entries) && !t.entries[i].isNil {
				node.Right = &TreeNode{Val: t.entries[i].value}
				next = append(next, node.Right)
			}
			i++
		}
		prev = &next
	}

	return root
}
