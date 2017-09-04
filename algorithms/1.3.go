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
	spacing := len(str) - length

	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[i+spacing] = '0'
			str[i+spacing-1] = '2'
			str[i+spacing-2] = '%'
			spacing -= 2
		} else {
			str[i+spacing] = str[i]
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
