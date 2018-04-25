// *Partition:* Write code to partition a linked list around value "x", such
// that all values less than "x" come before all values greater than or equal
// to "x".

package main

import (
	"container/list"
)

func partition(l list.List, x int) list.List {
	l2 := list.List{}

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) < x {
			l2.PushFront(e.Value)
		} else {
			l2.PushBack(e.Value)
		}
	}

	return l2
}
