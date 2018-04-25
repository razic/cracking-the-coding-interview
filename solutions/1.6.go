// *String Compression:* Implement a method to perform basic string compression
// using the counts of repeated characters. For example, the string
// "aabcccccaaa" would become "a2b1c5a3". If the compressed string would not
// become smaller than the original string, your method should return the
// original string.

package main

import (
	"bytes"
	"math"
	"strconv"
)

func compress(str string) string {
	indicies := []int{0}
	runes := bytes.Runes([]byte(str))

	// less than 3 runes, can't make it any smaller
	if len(runes) < 3 {
		return str
	}

	// count starting indicies of character sequences
	for i := 1; i < len(runes); i++ {
		if runes[i] != runes[i-1] {
			indicies = append(indicies, i)
		}
	}

	// inject the length at the end of the indicies for computation
	indicies = append(indicies, len(runes))

	// compute lengths of character sequences
	lengths := []int{}
	for i := 1; i < len(indicies); i++ {
		lengths = append(lengths, indicies[i]-indicies[i-1])
	}

	// compute the length of the would-be compressed string
	newLen := 0
	for i := 0; i < len(lengths); i++ {
		newLen += int(math.Log10(float64(lengths[i]))) + 2
	}

	// check to see if it still makes sense to perform the compression
	if newLen >= len(runes) {
		return str
	}

	// write out the bytes
	var buf bytes.Buffer
	for i := 0; i < len(lengths); i++ {
		buf.WriteRune(runes[indicies[i]])
		buf.WriteString(strconv.Itoa(lengths[i]))
	}

	return buf.String()
}
