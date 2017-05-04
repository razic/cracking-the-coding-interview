// *1.1 Is Unique:* Implement an algorithm to determine if a string has all
// unique characters. What if you cannot use additional data structures?

package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	str := os.Args[1]
	strlen := len(str)

	for i := 0; i < strlen; i++ {
		for j := 0; j < strlen; j++ {
			if i == j {
				continue
			}

			if str[i] == str[j] {
				os.Exit(1)
			}
		}
	}
}
