// ### Longest Common Subsequence
//
// Given two sequences, find the length of longest subsequence present in both of them. A subsequence
// is a sequence that appears in the same relative order, but not necessarily contiguous.
// For example, “abc”, “abg”, “bdf”, “aeg”, ‘”acefg”, .. etc are subsequences of “abcdefg”.
// So a string of length n has 2^n different possible subsequences.
//
// Examples:
// LCS for input Sequences “ABCDGH” and “AEDFHR” is “ADH” of length 3.
// LCS for input Sequences “AGGTAB” and “GXTXAYB” is “GTAB” of length 4.
//
// https://en.wikipedia.org/wiki/Hunt%E2%80%93McIlroy_algorithm
//
// http://introcs.cs.princeton.edu/java/96optimization/LCS.java.html
package main
import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lcs(a, b string) {
	M := len(a)
	N := len(b)

	// opt[i][j] = length of LCS of x[i..M] and y[j..N]
	opt := make([][]int, M+1)
	for i := 0; i < M+1; i++ {
		opt[i] = make([]int, N+1)
	}

	// compute length of LCS and all subproblems via dynamic programming
	for i := M-1; i >= 0; i-- {
		for j := N-1; j >= 0; j-- {
			if a[i] == b[j] {
				opt[i][j] = opt[i+1][j+1] + 1
			} else {
				opt[i][j] = max(opt[i+1][j], opt[i][j+1])
			}
		}
	}

	// recover LCS itself and print it to standard output
	i, j := 0, 0
	for i < M && j < N {
		if a[i] == b[j] {
			fmt.Print(string(a[i]))
			i++
			j++
		} else if opt[i+1][j] >= opt[i][j+1] {
			i++
		} else {
			j++
		}
	}
	fmt.Println()
}

func main() {
	lcs("ABCDGH", "AEDFHR")
}
// "ADH"

// https://en.wikipedia.org/wiki/Diff
//
// From a longest common subsequence it is only a small step to get diff-like output:
// if an item is absent in the subsequence but present in the first original sequence,
// it must have been deleted (as indicated by the '–' marks, below).
// If it is absent in the subsequence but present in the second original sequence,
// it must have been inserted (as indicated by the '+' marks).
