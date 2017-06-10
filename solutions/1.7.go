// *Rotate Matrix:* Given an image represented by a NxN matrix, where each
// pixel is 4 bytes, write a method to rotate the image by 90 degrees. Can you
// do this in place?

package main

import "fmt"

func rotate(matrix [][]int, n int) {
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			// top = left
			matrix[first][i] = matrix[last-offset][first]

			// left = bottom
			matrix[last-offset][first] = matrix[last][last-offset]

			// bottom = right
			matrix[last][last-offset] = matrix[i][last]

			// right =  top
			matrix[i][last] = top
		}
	}
}

func main() {
	m := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	for _, row := range m {
		fmt.Printf("%v\n", row)
	}

	fmt.Printf("\n")

	rotate(m, 4)

	for _, row := range m {
		fmt.Printf("%v\n", row)
	}
}
