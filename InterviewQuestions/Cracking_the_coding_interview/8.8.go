// Write an algorithm to print all ways of arranging eight queens on a chess board so that
// none of them share the same row, column or diagonal.
package main

import (
	"fmt"
)

const (
	Size = 8
)

func printBoard(cols []int) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if j == cols[i] {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func placeQueens2(n int, cols []int, ways *[][]int) {
	if n == Size {
		validCols := make([]int, len(cols))
		copy(validCols, cols)
		*ways = append(*ways, validCols)
		return
	}
	for i := 0; i < Size; i++ {
		cols[n] = i
		ok := true
		for j := 0; j < n; j++ {
			if cols[n] == cols[j] || n-j == cols[n]-cols[j] || n-j == cols[j]-cols[n] {
				ok = false
				break
			}
		}
		if ok {
			// if cols[0] == 0 {
			// 	fmt.Println("ok", n, cols)
			// }
			placeQueens2(n+1, cols, ways)
		}
	}
}

func main() {
	cols := make([]int, Size)
	ways := [][]int{}
	placeQueens(0, cols)
	fmt.Println("--------")
	cols = make([]int, Size)
	placeQueens2(0, cols, &ways)
	printBoard(ways[0])
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func placeQueens(row int, cols []int) {
	if row == Size {
		// found valid placement
		printBoard(cols)
	} else {
		for col := 0; col < Size; col++ {
			if checkValid(cols, row, col) {
				cols[row] = col
				placeQueens(row+1, cols)
			}
		}
	}
}

func checkValid(cols []int, row1, col1 int) bool {
	for row2 := 0; row2 < row1; row2++ {
		col2 := cols[row2]
		// check if (row2, col2) invalidates (row1, col1) as a queen spot

		// check if rows have a queen in same column
		if col1 == col2 {
			return false
		}

		// check diagonals: if the distance between the columsn equals the distance between
		// the rows, then they're in the same diagonal
		colDistance := Abs(col2 - col1)
		// row1 > row2, so no need for abs
		rowDistance := row1 - row2
		if colDistance == rowDistance {
			return false
		}
	}
	return true
}
