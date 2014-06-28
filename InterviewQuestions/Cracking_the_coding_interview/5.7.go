//
// An array A[1…n] contains all the integers from 0 to n except for one number which is missing.
// In this problem, we cannot access an entire integer in A with a single operation.
// The elements of A are represented in binary, and the only operation we can use
// to access them is “fetch the jth bit of A[i]”, which takes constant time.
// Write code to find the missing integer. Can you do it in O(n) time?
//

// If N is odd, count[1](0) > count[1](1)
// If N is even, count[1](0) == count[1](1)
// ie. N is odd and count[1](0) == count[1](1) => missing 0
//              and count[1](0) > count[1](1)  => missing 1
// ... count[2] ... count[N]

package main

import (
	"fmt"
	// "math"
)

func makeArray(n, missing int) []int {
	nums := make([]int, n-1)
	for i, j := 0, 0; i < n; i++ {
		if i != missing {
			nums[j] = i
			j++
		}
	}
	return nums
}

func findMissing(nums []int) int {
	return findMissingBit(nums, 0)
}

func findMissingBit(nums []int, j int) int {
	if j >= 64 {
		return 0
	}
	zeros, ones := counts(nums, j)
	if len(zeros) <= len(ones) {
		v := findMissingBit(zeros, j+1)
		return (v << 1) | 0
	} else {
		v := findMissingBit(ones, j+1)
		return (v << 1) | 1
	}
}

func fetch(nums []int, i, j int) int {
	return (nums[i] >> uint(j)) & 1
}

func counts(nums []int, j int) ([]int, []int) {
	ones := []int{}
	zeros := []int{}
	for i, _ := range nums {
		if fetch(nums, i, j) == 1 {
			ones = append(ones, nums[i])
		} else {
			zeros = append(zeros, nums[i])
		}
	}
	return zeros, ones
}

func main() {
	fmt.Println("5.7 - find missing number")
	nums := makeArray(11, 4)
	fmt.Println("test fetch", fetch(nums, 1, 0))
	missing := findMissing(nums)
	fmt.Printf("0 - %d, missing: %d, in binary %b\n", len(nums)+1, missing, missing)
	nums = makeArray(11, 5)
	missing = findMissing(nums)
	fmt.Printf("missing: %d, in binary %b\n", missing, missing)
	nums = makeArray(10, 5)
	missing = findMissing(nums)
	fmt.Printf("missing: %d, in binary %b\n", missing, missing)
	nums = makeArray(10, 8)
	missing = findMissing(nums)
	fmt.Printf("missing: %d, in binary %b\n", missing, missing)
}
