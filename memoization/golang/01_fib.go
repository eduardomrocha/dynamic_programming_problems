// Fibonacci.

// Implementation of the fibonacci function using memoization.

// The fibonacci function takes in a number as an argument.
// The function should return the n-th number of the Fibonacci sequence.
// The first and second number of the sequence is 1.
package main

import "fmt"

// Find the n-th number of the Fibonacci sequence.
func Fibonacci(n int, memo *map[int]int) int {
	if ans, ok := (*memo)[n]; ok {
		return ans
	}

	if n <= 2 {
		return 1
	}

	(*memo)[n] = Fibonacci(n-1, memo) + Fibonacci(n-2, memo)
	return (*memo)[n]
}

func main() {
	memo := make(map[int]int)

	fmt.Println(Fibonacci(6, &memo))  // 8
	fmt.Println(Fibonacci(7, &memo))  // 13
	fmt.Println(Fibonacci(8, &memo))  // 21
	fmt.Println(Fibonacci(50, &memo)) // 12586269025
}
