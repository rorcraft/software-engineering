//
// Explain what the following code does: ((n & (n-1)) == 0).
//
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("5.4 (n & (n-1) == 0)")
	num, _ := strconv.ParseInt("100000", 2, 64)
	fmt.Printf("n: %b, (n & (n-1) == 0): %b\n", num, (num & (num-1)))
	num, _ = strconv.ParseInt("110000", 2, 64)
	fmt.Printf("n: %b, (n & (n-1) == 0): %b\n", num, (num & (num-1)))
	num, _ = strconv.ParseInt("1111", 2, 64)
	fmt.Printf("n: %b, (n & (n-1) == 0): %b\n", num, (num & (num-1)))
	fmt.Println("checks whether n has all (n-1) trailing zeros")
}
