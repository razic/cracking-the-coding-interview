// *Delete Middle Node:* Implement an algorithm to delete a node in the middle
// of a singly linked list, given only access to that node.

package main

import (
	"container/list"
)

func deleteMiddleNode(e *list.Element) bool {
	if e == nil || e.Next() == nil {
		return false
	}

	// swap pointer of the element with the next one
	*e = *e.Next()

	return true
}
