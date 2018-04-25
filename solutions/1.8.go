// *Zero Matrix:* Write an algorithm such that if an element in an NxM matrix
// is 0, it's entire row and column are set to 0.

package main

func zeroMatrix(matrix [][]int, n, m int) {
	rows, cols := make(map[int]bool), make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		if rows[i] == true {
			matrix[i] = make([]int, m)
			continue
		}

		for j := 0; j < len(cols); j++ {
			matrix[i][j] = 0
		}
	}
}
