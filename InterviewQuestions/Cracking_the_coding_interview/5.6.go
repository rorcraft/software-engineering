// Write a program to swap odd and even bits in an integer with as few instructions
// as possible (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, etc).
package main

import (
	"fmt"
	"strconv"
)

func swapbits(num int64) int64 {
	return ((num & 0x55) << 1) | ((num & 0xAA) >> 1)
	// return ((num & 0x55555555) << 1) | ((num >> 1) & 0x55555555)
}

func main() {
	fmt.Println("5.6 - swap bits")
	fmt.Printf("0x55555555 = %b\n", 0x55)
	fmt.Printf("0xAAAAAAAA = %b\n", 0xAA)
	num, _ := strconv.ParseInt("111011", 2, 64)
	fmt.Printf("%b & %b >> 1 %b\n", num, 0x55, (num & (0x555)) )
	fmt.Printf("%b & %b >> 1 %b\n", num, 0xAA, (num & (0xAAA)) )
	fmt.Printf("%b swapped %b\n", num, swapbits(num))
	num, _ = strconv.ParseInt("101010", 2, 64)
	fmt.Printf("%b swapped %b\n", num, swapbits(num))
	num, _ = strconv.ParseInt("1101010", 2, 64)
	fmt.Printf("%b swapped\n%b\n", num, swapbits(num))
}
