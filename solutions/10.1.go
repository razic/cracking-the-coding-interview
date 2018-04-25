// *Sorted Merge:* You are given two sorted arrays, A and B, where A has a
// large enough buffer at the end to hold B. Write a method to merge B into A
// in sorted order.

package main

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
