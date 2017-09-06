package tree

type Tree interface {
	Root() Node
	SetRoot(Node)
}

type Node interface {
	Value() interface{}
	SetValue(interface{})
	Append(...Node)
	Children() []Node
}

type tree struct {
	root Node
}

type node struct {
	value    interface{}
	children []Node
}

func NewNode(v interface{}) Node {
	return &node{value: v}
}

func NewTree(root Node) Tree {
	return &tree{root: root}
}

func (t *tree) Root() Node {
	return t.root
}

func (t *tree) SetRoot(n Node) {
	t.root = n
}

func (n *node) Append(ns ...Node) {
	n.children = append(n.children, ns...)
}

func (n *node) Children() []Node {
	return n.children
}

func (n *node) Value() interface{} {
	return n.value
}

func (n *node) SetValue(v interface{}) {
	n.value = v
}
