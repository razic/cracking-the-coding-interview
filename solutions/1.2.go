// *Check Permutation:* Given two strings, write a method to decide if one is a
// permutation of the other.

package main

func CheckPermutation(strA, strB string) bool {
	// lengths are different, can't be a permutation
	if len(strA) != len(strB) {
		return false
	}

	// string is the same, can't be a permutation
	if strA == strB {
		return false
	}

	// make a map to store occurences of runes in both strings
	runeCounts := make(map[rune][]int)

	// initialize zero value for rune counts in first string
	for _, r := range strA {
		runeCounts[r] = []int{0, 0}
	}

	// initialize zero value for rune counts in second string
	for _, r := range strB {
		runeCounts[r] = []int{0, 0}
	}

	// count occurences of each rune in first string
	for _, r := range strA {
		runeCounts[r][0]++
	}

	// count occurences of each rune in second string
	for _, r := range strB {
		runeCounts[r][1]++
	}

	// validate the occurrences of each rune match up
	for _, counts := range runeCounts {
		if counts[0] != counts[1] {
			return false
		}
	}

	// if we got this far, it's a valid permutation
	return true
}
