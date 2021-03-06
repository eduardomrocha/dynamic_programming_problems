// Fibonacci.

// Implementation of the fibonacci function using tabulation.

// The fibonacci function takes in a number as an argument.
// The function should return the n-th number of the Fibonacci sequence.
// The 0th number of the sequence is 0 and the 1st number of the sequence is 1.
// To generate the next number of the sequence, we sum the previous two.
package main

import "fmt"

// Given a integer n, return the n-th number of the fibonacci sequence.
func Fibonacci(n int) int {
	var table []int

	for i := 0; i <= n; i++ {
		table = append(table, 0)
	}

	table[1] = 1

	for i := 0; i < len(table); i++ {
		if i+1 < len(table) {
			table[i+1] += table[i]
			if i+2 < len(table) {
				table[i+2] += table[i]
			}
		}
	}

	return table[n]
}

func main() {
	fmt.Println(Fibonacci(6))  // 8
	fmt.Println(Fibonacci(7))  // 13
	fmt.Println(Fibonacci(8))  // 21
	fmt.Println(Fibonacci(50)) // 12586269025
}
