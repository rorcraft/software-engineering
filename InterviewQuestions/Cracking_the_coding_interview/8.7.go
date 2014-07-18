// Given an infinite number of quarters (25 cents), dimes (10 cents), nickels (5 cents)
// and pennies (1 cent), write code to calculate the number of ways of representing n cents.
//
// http://hawstein.com/posts/8.7.html
package main

import "fmt"

// 25, 10, 5, 1
func makeChange(n, denom int) int {
	next_denom := 0
	switch denom {
	case 25:
		next_denom = 10
		break
	case 10:
		next_denom = 5
		break
	case 5:
		next_denom = 1
		break
	case 1:
		return 1
	}
	ways := 0
	for i := 0; i*denom <= n; i++ {
		ways += makeChange(n-i*denom, next_denom)
	}
	return ways
}

func main() {
	fmt.Println("8.7 ways of combination of changes")
	fmt.Println("25", makeChange(25, 25), "ways")
	fmt.Println("125", makeChange(125, 25), "ways")
	fmt.Println("15", makeChange(15, 25), "ways")
	fmt.Println("625", makeChange(625, 25), "ways")
}
