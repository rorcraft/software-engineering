//
// Write a function to determine the number of bits required to convert integer A to integer B.
//
// Input: 31, 14
//
// Output: 2
package main

import (
	"fmt"
	"strconv"
)

func diffbits(a, b int64) int {
	diff := a ^ b
	ones := 0
	for diff > 0 {
		if (diff & 1) == 1 {
			ones++
		}
		diff >>= 1
	}
	return ones
}
func main() {
	a, _ := strconv.ParseInt("10011", 2, 64)
	b, _ := strconv.ParseInt("10000", 2, 64)
	fmt.Printf("Diff %d, %d = %d\n", a, b, diffbits(a, b))
	a = 31
	b = 14
	fmt.Printf("Diff %d, %d = %d\n", a, b, diffbits(a, b))
}
