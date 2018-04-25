// *Palindrome:* Implement a function to check if a linked list is a
// palindrome.

package main

import (
	"container/list"
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
