// *Route Between Nodes:* Given a directed graph, design an algorithm to find
// out whether there is a route between two nodes.

package main

import (
	"fmt"

	"github.com/razic/cracking-the-coding-interview/datastructures/graph"
)

func main() {
	// init a graph
	g := graph.NewGraph()

	// init some points
	n0 := graph.NewNode(1)
	n1 := graph.NewNode(1)
	n2 := graph.NewNode(2)
	n3 := graph.NewNode(3)
	n4 := graph.NewNode(4)
	n5 := graph.NewNode(5)
	n6 := graph.NewNode(6)

	// place the points in the graph
	g.Append(n0, n1, n2, n3, n4, n5, n6)

	// make some connections between nodes
	n0.Append(n1, n2)
	n2.Append(n1, n3)
	n3.Append(n2, n4)
	n4.Append(n5, n6)
	n6.Append(n4)

	// check to see if there is a route between some nodes
	if route(n0, n6) != true {
		fmt.Printf("your algorithm is borked")
	}
	if route(n0, n2) != true {
		fmt.Printf("your algorithm is borked")
	}
	if route(n6, n3) != false {
		fmt.Printf("your algorithm is borked")
	}
}

func route(a, b graph.Node) bool {
	return bfs(a, b)
}

func bfs(a, b graph.Node) bool {
	visited := map[graph.Node]bool{} // needed to detect loops
	queue := []graph.Node{a}         // used to perform traversal

	for len(queue) > 0 {
		a = queue[0]
		queue = queue[1:] // pop the queue

		if a == b {
			return true // route found
		}

		// add children to queue
		for _, e := range a.Adjacent() {
			if visited[e] == false {
				queue = append(queue, e)
			}
		}

		visited[a] = true // mark as visited
	}

	// no route found
	return false
}
