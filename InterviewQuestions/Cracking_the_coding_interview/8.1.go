// Write a method to generate the nth Fibonacci number.
package main

import "fmt"

// fib(n) = fib(n-1) + fib(n-2)

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	p1, p2 := 1, 1
	for i := 3; i < n; i++ { //iterative
		p1, p2 = p2, p1+p2
	}
	return p1 + p2
}

func main() {
	fmt.Println("8.1 Fibonacci")
	fmt.Println("3", fibonacci(3))
	fmt.Println("4", fibonacci(4))
	fmt.Println("5", fibonacci(5))
	fmt.Println("6", fibonacci(6))
	fmt.Println("60", fibonacci(60))
}
