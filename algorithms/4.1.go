// *Route Between Nodes:* Given a directed graph, design an algorithm to find
// out whether there is a route between two nodes.

package main

import (
	"fmt"

	"github.com/razic/cracking-the-coding-interview/datastructures"
)

func main() {
	// init a graph
	graph := &datastructures.Graph{}

	// init some points
	node0 := &datastructures.GraphNode{Value: 0}
	node1 := &datastructures.GraphNode{Value: 1}
	node2 := &datastructures.GraphNode{Value: 2}
	node3 := &datastructures.GraphNode{Value: 3}
	node4 := &datastructures.GraphNode{Value: 4}
	node5 := &datastructures.GraphNode{Value: 5}
	node6 := &datastructures.GraphNode{Value: 6}

	// place the points in the graph
	graph.Nodes = append(graph.Nodes, node0)
	graph.Nodes = append(graph.Nodes, node1)
	graph.Nodes = append(graph.Nodes, node2)
	graph.Nodes = append(graph.Nodes, node3)
	graph.Nodes = append(graph.Nodes, node4)
	graph.Nodes = append(graph.Nodes, node5)
	graph.Nodes = append(graph.Nodes, node6)

	// make some connections between nodes
	node0.Adjacent = append(node0.Adjacent, node1)
	node0.Adjacent = append(node0.Adjacent, node2)
	node2.Adjacent = append(node2.Adjacent, node1)
	node2.Adjacent = append(node2.Adjacent, node3)
	node3.Adjacent = append(node3.Adjacent, node2)
	node3.Adjacent = append(node3.Adjacent, node4)
	node4.Adjacent = append(node4.Adjacent, node5)
	node4.Adjacent = append(node4.Adjacent, node6)
	node6.Adjacent = append(node6.Adjacent, node4)

	// check to see if there is a route between some nodes
	if route(node0, node6) != true {
		fmt.Printf("your algorithm is borked")
	}
	if route(node0, node2) != true {
		fmt.Printf("your algorithm is borked")
	}
	if route(node6, node3) != false {
		fmt.Printf("your algorithm is borked")
	}
}

func route(a, b *datastructures.GraphNode) bool {
	return bfs(a, b)
}

func bfs(a, b *datastructures.GraphNode) bool {
	visited := map[*datastructures.GraphNode]bool{}
	queue := []*datastructures.GraphNode{a}

	for len(queue) > 0 {
		a = queue[0]
		queue = queue[1:]

		if a == b {
			return true
		}

		for _, e := range a.Adjacent {
			if visited[e] == false {
				queue = append(queue, e)
			}
		}

		visited[a] = true
	}

	return false
}
