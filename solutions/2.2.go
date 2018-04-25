// *Return Kth to Last:* Implement an algorithm to return the kth to the last
// element of a singly linked list. Assume you don't know the length.

package main

import (
	"container/list"
)

func kthFromLast(k int, l list.List) *list.Element {
	a := l.Front()
	b := l.Front()
	mid := 0

	// use fast pointer technique to get midpoint
	for b.Next() != nil && b.Next().Next() != nil {
		a = a.Next()
		b = b.Next().Next()
		mid++
	}

	// out of bounds check
	if k > mid*2 {
		return nil
	}

	var c, d *list.Element

	// choose the correct starting point
	if k > mid {
		c = l.Front()
	} else {
		c = a
	}

	d = c

	// find the kth from last element
	for i := 0; i < mid-k; i++ {
		d = d.Next()
	}

	return d
}
