// *Partition:* Write code to partition a linked list around value "x", such
// that all values less than "x" come before all values greater than or equal
// to "x".

package main

import (
	"container/list"
	"fmt"
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

func main() {
	l := list.List{}

	l.PushBack(1)
	l.PushBack(8)
	l.PushBack(3)
	l.PushBack(7)
	l.PushBack(2)

	l2 := partition(l, 3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value)
	}

	fmt.Printf("\n")

	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value)
	}

	fmt.Printf("\n")
}
