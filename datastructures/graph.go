package datastructures

type GraphNode struct {
	Value    interface{}
	Adjacent []*GraphNode
}

type Graph struct {
	Nodes []*GraphNode
}
