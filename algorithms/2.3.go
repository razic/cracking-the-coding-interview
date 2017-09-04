// *Delete Middle Node:* Implement an algorithm to delete a node in the middle
// of a singly linked list, given only access to that node.

package main

import (
	"container/list"
	"fmt"
)

func deleteMiddleNode(e *list.Element) bool {
	if e == nil || e.Next() == nil {
		return false
	}

	// swap pointer of the element with the next one
	*e = *e.Next()

	return true
}

func main() {
	l := list.List{}

	l.PushBack(1)
	l.PushBack(2)
	e := l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value)
	}

	fmt.Printf("\n")
	deleteMiddleNode(e)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value)
	}

	fmt.Printf("\n")
}
