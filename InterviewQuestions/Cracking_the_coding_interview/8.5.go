// Implement an algorithm to print all valid (e.g., properly opened and closed) combinations
// of n-pairs of parentheses.
//
// EXAMPLE:
//
// input: 3 (e.g., 3 pairs of parentheses)
//
// output: ((())), (()()), (())(), ()(()), ()()()

// like dfs ? global n-1, traverse down
package main

import "fmt"

func ParensCombo(n int) []string {
	return parensRecur(n, 0)
}

func parensRecur(n, opened int) []string {
	var combos []string
	if n == 0 {
		str := ""
		for opened > 0 {
			str += ")"
			opened--
		}
		return []string{str}
	}
  if opened > 0 {
		for _, rest := range parensRecur(n, opened-1) {
			combos = append(combos, ")" + rest)
		}
	}
	for _, rest := range parensRecur(n-1, opened+1) {
		combos = append(combos, "(" + rest)
	}
	return combos
}

func main() {
	fmt.Println("8.5 - Print all valid combinations of n-pairs of parentheses")
	fmt.Println(3, ParensCombo(3))
	fmt.Println(5, ParensCombo(5))
}
