package graph

type Graph interface {
	Append(...Node)
	Nodes() []Node
}

type graph struct {
	nodes []Node
}

type Node interface {
	Append(...Node)
	Adjacent() []Node
	Value() interface{}
}

type node struct {
	adjacent []Node
	value    interface{}
}

func NewGraph() Graph {
	return &graph{}
}

func NewNode(v interface{}) Node {
	return &node{value: v}
}

func (g *graph) Append(ns ...Node) {
	g.nodes = append(g.nodes, ns...)
}

func (g *graph) Nodes() []Node {
	return g.nodes
}

func (n *node) Append(ns ...Node) {
	n.adjacent = append(n.adjacent, ns...)
}

func (n *node) Adjacent() []Node {
	return n.adjacent
}

func (n *node) Value() interface{} {
	return n.value
}
