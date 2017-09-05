// *Sum Lists:* You have two numbers represented by a linked list, where each
// node contains a single digit. The digits are stored in reverse order, such
// that the 1's digit is at the head of the list. Write a function that adds
// the two numbers and returns the sum as a linked list.

package main

import (
	"container/list"
	"fmt"
)

func listToInt(l list.List) int {
	n, p := 0, 1

	for e := l.Front(); e != nil; e = e.Next() {
		n += (e.Value.(int) * p)
		p *= 10
	}

	return n
}

func intToList(n int) list.List {
	digits := list.List{}

	for ; n > 0; n /= 10 {
		digits.PushBack(n % 10)
	}

	return digits
}

func sumLists(a, b list.List) list.List {
	return intToList(listToInt(a) + listToInt(b))
}

func main() {
	a, b := list.List{}, list.List{}

	a.PushBack(7)
	a.PushBack(1)
	a.PushBack(6)
	b.PushBack(5)
	b.PushBack(9)
	b.PushBack(2)

	c := sumLists(a, b)

	for e := c.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value)
	}
}
