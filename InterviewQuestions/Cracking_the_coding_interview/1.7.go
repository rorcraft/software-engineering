package main
/*
 * Write an algorithm such that if an element in an MxN matrix is 0,
 * its entire row and column is set to 0.
 */
import (
	"fmt"
)
type Matrix [][]int32

func zeroRowCol(matrix Matrix) Matrix {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[0] {
			if col[j] || row[i] {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func printMatrix(matrix Matrix) {
	for i := range matrix {
		for j := range matrix[0] {
			fmt.Printf("%v\t", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	matrix := Matrix{
						[]int32{1, 2, 3},
						[]int32{1, 2, 0},
						[]int32{3, 3, 3},
						[]int32{0, 1, 1}}

	printMatrix(matrix)
	fmt.Println("zero rows and cols")
	printMatrix(zeroRowCol(matrix))
}
