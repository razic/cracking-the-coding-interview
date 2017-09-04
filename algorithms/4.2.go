// *Minimal Tree:* Given a sorted (ascending order) array with unique integer
// elements, write an algorithm to create a binary search tree with minimal
// height.

package main

import (
	"fmt"

	"github.com/razic/cracking-the-coding-interview/datastructures"
)

func main() {
	ints := []int{0, 4, 6, 9, 11, 16, 17, 100}
	tree := makeTree(ints, 2)

	// pre-order breadth-first search on a search tree should visit in order
	bfs(tree.Nodes[0], func(n *datastructures.GraphNode) {
		fmt.Printf("%d\n", n.Value)
	})
}

// expects array of ints to be in a specific order
func makeTree(ints []int, ary int) *datastructures.Graph {
	// init a tree
	tree := &datastructures.Graph{}
	// init the root node
	root := &datastructures.GraphNode{Value: ints[0]}
	// pop ints
	ints = ints[1:]

	// bfs
	bfs(root, func(n *datastructures.GraphNode) {
		// get max of len(ints) or ary (for incomplete trees)
		if len(ints) < ary {
			ary = len(ints)
		}

		// populate the adjacent nodes from the int list
		for _, i := range ints[:ary] {
			n.Adjacent = append(n.Adjacent, &datastructures.GraphNode{Value: i})
		}

		// move the int list
		ints = ints[ary:]
	})

	tree.Nodes = append(tree.Nodes, root)

	return tree
}

func bfs(a *datastructures.GraphNode, visit func(n *datastructures.GraphNode)) {
	queue := []*datastructures.GraphNode{a}
	for len(queue) > 0 {
		a = queue[0]
		queue = queue[1:]
		visit(a)
		queue = append(queue, a.Adjacent...)
	}
}
