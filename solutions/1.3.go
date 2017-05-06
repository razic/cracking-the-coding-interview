// *URLify:* Write a method to replace all spaces in a string with '%20'. You
// may assume that the string has sufficient space at the end to hold
// additional characters, and that you are given the "true" length of the
// string.
//
// Example
//  Input: "Mr John Smith    "
//  Output: "Mr%20John%20Smith"
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	str := []byte(os.Args[1])
	truelen, err := strconv.Atoi(os.Args[2])

	if err != nil {
		os.Exit(1)
	}

	// trim string from the end to be the true length
	str = str[:truelen]

	// count actual whitespaces and save their indicies
	spaces := []int{}
	for i, n := range str {
		if n == ' ' {
			spaces = append(spaces, i)
		}
	}

	for i, n := range spaces {
		rem := make([]byte, truelen-n)

		// make a copy of everything after the space
		copy(rem, str[n+(i*2)+1:])

		// slice the original from the beginning to the space
		str = str[:n+(i*2)]

		str = append(str, []byte("%20")...)
		str = append(str, rem...)
	}

	fmt.Printf("%s\n", str)
}
