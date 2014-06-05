package main

import (
	"fmt"
)
/*
 * Question
 * Implement an algorithm to determine if a string has all unique characters.
 * What if you can not use additional data structures?
 *
 * http://hawstein.com/posts/1.1.html
 */

/*
 * Notes
 * - ascii ? 256, a-z? upper case? utf8?
 */

// ascii only, using bitset
func isUnique(content string) bool {
	bitset := make([]uint64, 256 / 64)
	for _, c := range content {
		if bitset[c/64] & (1 << (uint(c)%64)) != 0 {
			return false
		}
		bitset[c/64] |= (1 << (uint(c)%64))
	}
	return true
}

// not using data structure O(n^2)
func isUnique1(content string) bool {
	for i, n := range content {
		for j, m := range content {
			if i != j && n == m {
				return false
			}
		}
	}
	return true
}

func main() {
	uniqStr := "abcdefghijkl"
	notUniqStr := "abcdefgggggg"
	fmt.Println(uniqStr, "is unique?", isUnique(uniqStr))
	fmt.Println(notUniqStr, "is unique?", isUnique(notUniqStr))
	fmt.Println(uniqStr, "is unique?", isUnique1(uniqStr))
	fmt.Println(notUniqStr, "is unique?", isUnique1(notUniqStr))
}
