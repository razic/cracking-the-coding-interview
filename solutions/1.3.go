// *URLify:* Write a method to replace all spaces in a string with '%20'. You
// may assume that the string has sufficient space at the end to hold
// additional characters, and that you are given the "true" length of the
// string.
//
// Example
//  Input: "Mr John Smith    ", 13
//  Output: "Mr%20John%20Smith"

package main

import (
	"fmt"
	"os"
	"strconv"
)

func urlify(str []byte, length int) {
	var spaceCount, newLength int

	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	newLength = length + (spaceCount * 2)

	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[newLength-1] = '0'
			str[newLength-2] = '2'
			str[newLength-3] = '%'
			newLength -= 3
		} else {
			str[newLength-1] = str[i]
			newLength--
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	str := []byte(os.Args[1])
	n, err := strconv.Atoi(os.Args[2])

	if err != nil {
		os.Exit(1)
	}

	urlify(str, n)

	fmt.Printf("%s\n", str)
}
