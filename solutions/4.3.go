// *List of Depths:* Given a binary tree, design an algorithm which creates a
// linked list of all the nodes at each depth.

package main

import (
	"container/list"
	"fmt"

	"github.com/razic/cracking-the-coding-interview/datastructures/tree"
)

func main() {
	tr := tree.NewBinaryTree(nil)
	n0 := tree.NewBinaryNode(2)
	n1 := tree.NewBinaryNode(3)
	n2 := tree.NewBinaryNode(4)
	n3 := tree.NewBinaryNode(5)
	n4 := tree.NewBinaryNode(6)
	n5 := tree.NewBinaryNode(7)
	n6 := tree.NewBinaryNode(8)
	n7 := tree.NewBinaryNode(9)

	n0.SetLeft(n1)
	n0.SetRight(n2)
	n1.SetLeft(n3)
	n2.SetLeft(n4)
	n3.SetLeft(n5)
	n3.SetRight(n6)
	n6.SetRight(n7)
	tr.SetRoot(n0)

	for i, v := range bfs(tr) {
		fmt.Printf("%d:", i)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf(" %d", e.Value)
		}
		fmt.Printf("\n")
	}
}

type qObj struct {
	n tree.BinaryNode
	i int
}

func bfs(t tree.BinaryTree) []*list.List {
	q := []*qObj{&qObj{n: t.Root(), i: 0}}
	count := 0
	levels := map[tree.BinaryNode]int{}

	for len(q) > 0 {
		qob := q[0]
		levels[qob.n] = qob.i
		q = q[1:]
		if qob.n.Left() != nil {
			q = append(q, &qObj{n: qob.n.Left(), i: qob.i + 1})
		}
		if qob.n.Right() != nil {
			q = append(q, &qObj{n: qob.n.Right(), i: qob.i + 1})
		}
		count++
	}

	// get the depth of the tree
	max := 0
	for _, v := range levels {
		if v > max {
			max = v
		}
	}

	// allocate lists for each level
	lists := make([]*list.List, max+1)
	for x := 0; x < max+1; x++ {
		lists[x] = list.New()
	}

	// bucket appropriately
	for k, v := range levels {
		lists[v].PushBack(k)
	}

	return lists
}
