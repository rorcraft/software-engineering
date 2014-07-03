// Write a method that returns all subsets of a set.
package main

import "fmt"

// AKA - power set - O(2^n)

// better implementation using interface - http://rosettacode.org/wiki/Power_set#Go

// not including empty set
func powerset(slice []int) [][]int {
	sets := [][]int{}
	for _, e := range slice {
		subsets := [][]int{}
		subsets = append(subsets, []int{e})
		// fmt.Println(e)
		for _, partialSet := range sets {
			// fmt.Println(append(partialSet, e))
			subsets = append(subsets, append(partialSet, e))
		}
		sets = append(sets, subsets...)
	}
	return sets
}

func main() {
	fmt.Println("8.3 - Powerset")
	set := []int{1, 2, 3, 4, 5}
	fmt.Println(set, powerset(set))
	// set = []int{1,2,3,4}
	// fmt.Println(set, powerset(set))
}
