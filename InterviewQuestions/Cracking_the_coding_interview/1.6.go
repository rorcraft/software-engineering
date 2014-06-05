/* Question
 * Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
 * write a method to rotate the image by 90 degrees. Can you do this in place?
 * 
 * http://hawstein.com/posts/1.6.html
 */

package main

import (
	"fmt"
	// "runtime"
)

func rotate(matrix [][]int32) [][]int32 {
	n := len(matrix)
	for i := range matrix {
		for j := i+1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n/2; i++ {
		for j := range matrix[0] {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	return matrix
}

func printMatrix(matrix [][]int32) {
	for i := range matrix {
		for j := range matrix[0] {
			fmt.Printf("%v\t", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		trace := make([]byte, 1024)
	// 		count := runtime.Stack(trace, true)
	// 		fmt.Printf("Recover from panic: %s\n", err)
	// 		fmt.Printf("Stack of %d bytes: %s\n", count, trace)
	// 	}
	// }()

	m := make([][]int32, 4)
	n := int32(1)
	for i := 0; i < 4; i++ {
		m[i] = make([]int32, 4)
		for j := 0; j < 4; j++ {
			m[i][j] = n
			n++
		}
	}
	printMatrix(m)
	fmt.Println("rotate")
	printMatrix(rotate(m))

} 
// $ go run 1.6.go
// 
// 1	2	3	4
// 5	6	7	8
// 9	10	11	12
// 13	14	15	16
// rotate
// 4	8	12	16
// 3	7	11	15
// 2	6	10	14
// 1	5	9	13
