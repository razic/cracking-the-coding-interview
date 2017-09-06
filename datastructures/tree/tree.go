package tree

type Tree interface {
	Root() Node
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

func NewTree(v interface{}) Tree {
	return &tree{root: NewNode(v)}
}

func (t *tree) Root() Node {
	return t.root
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
