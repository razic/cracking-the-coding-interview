// *Is Unique:* Implement an algorithm to determine if a string has all
// unique characters. What if you cannot use additional data structures?

package main

import (
	"os"
)

func isUnique(str string) bool {
	for i, runeA := range str {
		for j, runeB := range str {
			// skip if we are on the same index in both loops
			if i == j {
				continue
			}

			// we have a match, so the string is not unique
			if runeA == runeB {
				return false
			}
		}
	}

	return true
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	if !isUnique(os.Args[1]) {
		os.Exit(1)
	}
}
