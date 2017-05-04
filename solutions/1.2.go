// *Check Permutation:* Given two strings, write a method to decide if one is a
// permutation of the other.

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
		os.Exit(1)
	}

	n := len(a)

	if n != len(b) {
		os.Exit(1)
	}

	aocc := make(map[byte]int)
	bocc := make(map[byte]int)

	for i := 0; i < n; i++ {
		aocc[a[i]]++
		bocc[b[i]]++
	}

	for i := 0; i < n; i++ {
		if aocc[a[i]] != bocc[a[i]] {
			os.Exit(1)
		}
	}
}
