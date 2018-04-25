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
	"bytes"
)

func isOneAway(a, b string) bool {
	if a == b {
		return true
	}

	var (
		bufA, bufB      bytes.Buffer
		longer, shorter []rune
	)

	bufA.WriteString(a)
	bufB.WriteString(b)

	// capture unicode code points
	runesA := bytes.Runes(bufA.Bytes())
	runesB := bytes.Runes(bufB.Bytes())

	// capture the longer and shorter
	if len(runesA) > len(runesB) {
		longer = runesA
		shorter = runesB
	} else {
		longer = runesB
		shorter = runesA
	}

	lenDiff := len(longer) - len(shorter)
	discrepancies := 0

	// larger than a difference of 1 in length implies more than 1 edit
	if lenDiff > 1 {
		return false
	}

	for i, r := range shorter {
		// we have a match, so continue
		if r == longer[i] {
			continue
		}

		// if the length of the strings differ by 1, the next character over
		// should be our match, otherwise we've already got two discrepancies
		if lenDiff != 0 && r == longer[i+1] {
			continue
		}

		discrepancies++
	}

	// total up our discrepancies, plus the difference in length
	return discrepancies+lenDiff < 2
}
