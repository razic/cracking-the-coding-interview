// *Check Permutation:* Given two strings, write a method to decide if one is a
// permutation of the other.

package main

import (
	"os"
)

func checkPermutation(strA, strB string) bool {
	// if the length isn't the same, it can't be a permutation
	if len(strA) != len(strB) {
		os.Exit(1)
	}

	same := true
	occurrences := map[byte][2]int{}

	// in this loop we'll count the occurences for each byte in both strings
	// we'll also determine if the two strings are the same set of bytes
	for i := 0; i < len(strA); i++ {
		a := strA[i]
		b := strB[i]
		occa := occurrences[a][0]
		occb := occurrences[b][1]

		// the bytes don't match, they can't be the same string
		if a != b {
			same = false
		}

		occa++
		occb++
	}

	// same string, so not a permutation
	if same == true {
		return false
	}

	for _, v := range occurrences {
		if v[0] != v[1] {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	if !checkPermutation(os.Args[1], os.Args[2]) {
		os.Exit(1)
	}
}
