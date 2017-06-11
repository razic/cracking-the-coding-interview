// *String Rotation:* Assume you have a method which checks if one word is a
// substring of another. Given two strings, write code to check if one is a
// rotation of the other, using only a single call to the substring method.

package main

import (
	"bytes"
	"os"
	"strings"
)

func isRotation(a, b string) bool {
	var buf bytes.Buffer

	buf.WriteString(b)
	buf.WriteString(b)

	if strings.Contains(buf.String(), a) {
		return true
	}

	return false
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	if !isRotation(os.Args[1], os.Args[2]) {
		os.Exit(1)
	}
}
