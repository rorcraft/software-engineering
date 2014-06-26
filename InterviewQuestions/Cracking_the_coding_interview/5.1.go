/*
 * You are given two 32-bit numbers, N and M, and two bit positions, i and j.
 * Write a method to set all bits between i and j in N equal to M
 * (e.g., M becomes a substring of N located at i and starting at j).
 *
 * EXAMPLE:
 * Input: N = 10000000000, M = 10101, i = 2, j = 6
 *
 * Output: N = 10001010100
 */

package main

import (
	"fmt"
	"strconv"
)

func updateBits(n, m string, i, j uint) int64 {
	nint, _ := strconv.ParseInt(n, 2, 64)
	mint, _ := strconv.ParseInt(m, 2, 64)
	tail := ((1 << i) -1) & nint
	j = j+1
	head := ((nint >> j) << j)
	// fmt.Printf("head:\t%b\n", head)
	// fmt.Printf("n:\t%b\n", nint)
	return head | (mint << i) | tail
}

func main() {
	n := "10000000000"
	m := "10101"
	fmt.Println("N", n)
	fmt.Println("M", m)
	fmt.Printf("Update bits 2, 6:\t%b\n", updateBits(n, m, 2, 6))
	fmt.Printf("Update bits 6, 10:\t%b\n", updateBits(n, m, 6, 10))
}
