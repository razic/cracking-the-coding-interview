// *One Away:* There are three types of edits that can be performed on strings:
// insert a character, remove a character, or replace a character. Given two
// strings, write a function to check if they are one edit (or zero edits)
// away.
//
// Example:
//   pale,  ple  -> true
//   pales, pale -> true
//   pale,  bale -> true
//   pale,  bake -> false

package main

import (
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	a := os.Args[1]
	b := os.Args[2]

	if a == b {
		os.Exit(0)
	}

	diff := len(a) - len(b)

	// larger than a difference of 1 in length implies more than 1 edit
	if diff*diff > 1 {
		os.Exit(1)
	}

	if isOneAway(a, b, diff) == true {
		os.Exit(0)
	}

	os.Exit(1)
}

func isOneAway(a, b string, diff int) bool {
	var longer string
	var shorter string

	if len(a) > len(b) {
		longer = a
		shorter = b
	} else {
		longer = b
		shorter = a
	}

	discrepancies := 0

	for i := 0; i < len(shorter); i++ {
		// we have a match, so continue
		if shorter[i] == longer[i] {
			continue
		}

		// if the length of the strings differ by 1, the next character over
		// should be our match, otherwise we've already got two discrepancies
		if diff != 0 && shorter[i] == longer[i+1] {
			continue
		}

		discrepancies++
	}

	// total up our discrepancies, plus the absolute value of our difference
	// note: this squaring to obtain an absolute difference only works for the
	// value "1" and "-1". if this algorithm needed to detect more than one
	// edit, this, and the remainder of the algorithm would need to be modified
	// accordingly
	return discrepancies+(diff*diff) < 2
}
