package main

import "os"

func isPermutationOfPalindrome(str string) bool {
	pairs := make(map[rune]int)
	numRunes := 0
	numPairs := 0
	numSpaces := 0
	numPairsExpected := 0

	for _, r := range str {
		numRunes++

		// count the spaces and continue
		if r == ' ' {
			numSpaces++
			continue
		}

		pairs[r]++

		// found a pair
		if pairs[r] == 2 {
			pairs[r] = 0
			numPairs++
		}
	}

	// if there are less than 3 runes (minus spaces), can't be a palindrome
	if numRunes-numSpaces < 3 {
		return false
	}

	// calculate the number of expected pairs
	if (numRunes-numSpaces)%2 == 1 {
		numPairsExpected = ((numRunes - 1) - numSpaces) / 2
	} else {
		numPairsExpected = (numRunes - numSpaces) / 2
	}

	// if the number of pairs found matches whats expected, we have a true case
	if numPairs == numPairsExpected {
		return true
	}

	return false
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	if !isPermutationOfPalindrome(os.Args[1]) {
		os.Exit(1)
	}
}
