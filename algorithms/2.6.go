// *Palindrome:* Implement a function to check if a linked list is a
// palindrome.

package main

import (
	"container/list"
	"os"
)

func isPalindrome(l list.List) bool {
	a, b := l.Front(), l.Back()

	for i := 0; i < l.Len()/2; i++ {
		if a.Value != b.Value {
			return false
		}
	}

	return true
}

func main() {
	l := list.List{}

	l.PushBack("t")
	l.PushBack("a")
	l.PushBack("c")
	l.PushBack("o")
	l.PushBack("c")
	l.PushBack("a")
	l.PushBack("t")

	if !isPalindrome(l) {
		os.Exit(1)
	}
}
