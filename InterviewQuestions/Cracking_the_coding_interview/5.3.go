/*
 * Given an integer, print the next smallest and next largest number that have the
 * same number of 1 bits in their binary representation.
 */

// next largest = (least 1 digit > 0) >> 1
// next smallest = (least 1 digit << 1) + carry

package main

import (
	"fmt"
	"strconv"
)

func nextSmallest(num int64) int64 {
	// flip first 0
	// clear right, fill with same number of 1s
	i := uint(0)
	right := int64(0)
	for bits, one := num, false; bits > 0; bits >>= 1 {
		if (bits & 1) == 1 {
			right = (right << 1) | 1
			if !one {
				one = true
			}
		} else if one && (bits&1) == 0 {
			break
		}
		i++
	}
	left := (num>>i | 1) << i // set to 1, clear right
	right >>= 1
	return left | right
}

// previous int with same # of 1s
func nextLargest(num int64) int64 {
	// flip first 1
	// flip LS 0 bit
	i := uint(1)
	ones := uint(0)
	zero := false
	for bits := num >> 1; bits > 0; bits >>= 1 {
		if !zero && (bits&1) == 0 {
			zero = true
		}
		if (bits & 1) == 1 {
			ones++
			if zero {
				break
			}
		}
		i++
		// fmt.Printf("right %b\n", right)
	}
	if !zero {
		return num
	}
	left := (num >> (i + 1)) << (i + 1)
	right := (1<<(ones+1) - 1) << (i - ones - 1)
	// fmt.Printf("left %b\n", left)
	return left | int64(right)
}

func main() {
	fmt.Println("5.3 find next and previous binary with same number of 1s")
	num, _ := strconv.ParseInt("1111", 2, 64)
	fmt.Printf("num: %b\n", num)
	small := nextSmallest(num)
	fmt.Printf("next smallest %b\n", small)
	fmt.Printf("next largest %b\n", nextLargest(small))
	num, _ = strconv.ParseInt("100110", 2, 64)
	fmt.Printf("num: %b\n", num)
	small = nextSmallest(num)
	fmt.Printf("next smallest %b\n", small)
	fmt.Printf("next largest %b\n", nextLargest(small))
	num, _ = strconv.ParseInt("1111", 2, 64)
	fmt.Printf("num: %b\n", num)
	fmt.Printf("next largest %b\n", nextLargest(num))
}
