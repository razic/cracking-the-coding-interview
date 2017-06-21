// *Sorted Merge:* You are given two sorted arrays, A and B, where A has a
// large enough buffer at the end to hold B. Write a method to merge B into A
// in sorted order.

package main

import "fmt"

func merge(a, b []int) {
	indexA := len(a) - 1
	indexB := len(b) - 1
	indexC := cap(a) - 1

	// grow slice to cap
	a = a[:cap(a)]

	for indexB >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexC] = a[indexA]
			indexA--
		} else {
			a[indexC] = b[indexB]
			indexB--
		}
		indexC--
	}
}

func main() {
	a := make([]int, 3, 7)
	b := make([]int, 4)

	a[0] = 1
	a[1] = 5
	a[2] = 10
	b[0] = 1
	b[1] = 6
	b[2] = 9
	b[3] = 100

	merge(a, b)

	// grow slice to cap
	a = a[:cap(a)]

	fmt.Printf("%v\n", a)
}
