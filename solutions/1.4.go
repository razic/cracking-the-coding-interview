// *Palindrome Permutation:* Given a string, write a function to check if it is
// a permutation of a palindrome. A palindrome is a word or phrase that is the
// same forwards and backwards. A permutation is a rearrangement of letters.
// The palindrome does not need to be limited to dictionary words.
//
// Example:
//   Input: Tact Coa
//   Output: True (permutations: "taco cat", "atco cta", etc...)

package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	str := []byte(os.Args[1])

	if isPermutationOfPalindrome(str) {
		os.Exit(0)
	}

	os.Exit(1)
}

func isPermutationOfPalindrome(str []byte) bool {
	oddCount := 0
	counts := map[byte]int{}

	for _, c := range str {
		// we only care about ascii alphabet characters
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			// uppercase should be considered the same as its lowercase
			if c > 'Z' {
				c = c - 32
			}

			counts[c]++

			if counts[c]%2 == 1 {
				oddCount++
			} else {
				oddCount--
			}
		}
	}

	return oddCount <= 1
}
