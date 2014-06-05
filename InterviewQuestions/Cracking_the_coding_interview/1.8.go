/* Question
 * Assume you have a method isSubstring which checks if one word is a substring of another.
 * Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only
 * one call to isSubstring ( i.e., “waterbottle” is a rotation of “erbottlewat”).
 */

package main

import (
	"strings"
	"fmt"
)

func isSubstring(sub, str string) bool {
	return strings.Index(str, sub) >= 0
}

func isRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	str1 = str1 + str1
	return isSubstring(str2, str1)
}

func main() {
	fmt.Println(isSubstring("def", "abcdef"))
	fmt.Println(isSubstring("cde", "abcdef"))
	fmt.Println(isSubstring("fed", "abcdef"))
	str1, str2 := "waterbottle", "erbottlewat"
	fmt.Println(str1, str2, "is rotation?", isRotation(str1, str2))
	str1, str2 = "apple", "leape"
	fmt.Println(str1, str2, "is rotation?", isRotation(str1, str2))
	str1, str2 = "apple", "leapp"
	fmt.Println(str1, str2, "is rotation?", isRotation(str1, str2))
}
