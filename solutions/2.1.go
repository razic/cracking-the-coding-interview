// *Remove Dups:* Write code to remove duplicates from an unsorted linked list.

package main

import (
	"container/list"
)

func removeDups(l list.List) list.List {
	// create a new empty list
	l2 := list.List{}

	// create an empty hashmap for deduping values
	vals := make(map[interface{}]bool)

	// loop through the given list, storing element values in the hashmap
	for e := l.Front(); e != nil; e = e.Next() {
		vals[e.Value] = true
	}

	// loop through the deduped values, push into new list
	for v := range vals {
		l2.PushBack(v)
	}

	// return new list
	return l2
}
