// *Return Kth to Last:* Implement an algorithm to return the kth to the last
// element of a singly linked list. Assume you don't know the length.

package main

import (
	"container/list"
	"os"
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

func main() {
	l := list.List{}

	l.PushBack(8)
	l.PushBack(2)
	l.PushBack(9)
	l.PushBack(1)
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(7)

	i := 0

	for a, b := l.Back(), kthFromLast(0, l); a == b; i++ {
		a = a.Prev()
		b = kthFromLast(i, l)
	}

	if i == 0 {
		os.Exit(1)
	}
}
