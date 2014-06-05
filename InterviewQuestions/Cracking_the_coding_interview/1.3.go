/*
 * Question
 * Given two strings, write a method to decide if one is a permutation of another
 */

package main
import (
	"fmt"
)

func isPermu(one, another string) bool {
	oneCount, anotherCount := letterCount(one), letterCount(another)
	for i := 0; i < 256; i++ {
		if oneCount[i] != anotherCount[i] {
			return false
		}
	}
	return true
}

func letterCount(str string) []int {
	counts := make([]int, 256)
	for _, c := range str {
		counts[c]++
	}
	return counts
}

func main() {
	dog, god := "dog", "god"
	doc := "doc"
	fmt.Println(dog, god, "is permutation?", isPermu(dog, god))
	fmt.Println(dog, doc, "is permutation?", isPermu(dog, doc))
}

// Another solution is to sort the strings first and compare.
