// Write a method to compute all permutations of a string
package main

import "fmt"

// http://hawstein.com/posts/8.4.html
func permutations(str string) []string {
	perms := []string{}
	if len(str) == 0 {
		return append(perms, "")
	}
	for i, char := range str {
		var rest string
		rest = str[:i] + str[i+1:]
		// fmt.Println("char", string(char))
		// fmt.Println("rest", rest)
		for _, restPerm := range permutations(rest) {
			perms = append(perms, string(char) + restPerm)
		}
	}
	return perms
}

func main() {
	fmt.Println("8.4 - Permutations of string")
	str := "abcd"
	// str := "a"
	fmt.Println(str, permutations(str))
}
