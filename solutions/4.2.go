// *Minimal Tree:* Given a sorted (ascending order) array with unique integer
// elements, write an algorithm to create a binary search tree with minimal
// height.

package main

import (
	"fmt"

	"github.com/razic/cracking-the-coding-interview/datastructures/tree"
)

func main() {
	bfs(makeTree([]int{1, 3, 5, 9, 12, 24, 25}))
}

func makeTree(ints []int) tree.BinaryTree {
	rt := tree.NewBinaryNode(nil)
	tr := tree.NewBinaryTree(rt)

	fill(rt, ints)

	return tr
}

// recursively fill a tree
func fill(parent tree.BinaryNode, ints []int) {
	l := len(ints)

	if l == 0 {
		return
	}

	h := l / 2
	mid := ints[h]
	left := ints[:h]
	right := ints[h+1:]

	parent.SetValue(mid)
	parent.SetLeft(tree.NewBinaryNode(nil))
	parent.SetRight(tree.NewBinaryNode(nil))
	fill(parent.Left(), left)
	fill(parent.Right(), right)
}

// used to verify the ordering
func bfs(t tree.BinaryTree) {
	q := []tree.BinaryNode{t.Root()}

	for len(q) > 0 {
		n := q[0]
		if n.Value() != nil {
			fmt.Println(n.Value())
		}
		if n.Left() != nil {
			q = append(q, n.Left())
		}
		if n.Right() != nil {
			q = append(q, n.Right())
		}
		q = q[1:]
	}
}
