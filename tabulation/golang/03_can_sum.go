// Can sum.

// Implementation of a solution to the can sum problem using memoization.

// Given an integer targetSum and a list of numbers as arguments, return a boolean
// indicating whether or not it is possible to generate the targetSum using numbers from
// the list, using an element of the list as many times as needed and assuming that all
// input numbers are nonnegative.
package main

import "fmt"

// Check if the targetSum can be gereteded by adding up elements from numbers.
func CanSum(targetSum int, numbers []int) bool {
	var table []bool

	for i := 0; i <= targetSum; i++ {
		table = append(table, false)
	}

	table[0] = true

	for i := 0; i <= targetSum; i++ {
		if table[i] {
			for _, number := range numbers {
				if i+number <= targetSum {
					table[i+number] = true
				}
			}
		}
	}

	return table[targetSum]
}
func main() {
	fmt.Println(CanSum(7, []int{2, 3}))       // True
	fmt.Println(CanSum(7, []int{5, 3, 4, 7})) // True
	fmt.Println(CanSum(7, []int{2, 4}))       // False
	fmt.Println(CanSum(8, []int{2, 3, 5}))    // True
	fmt.Println(CanSum(300, []int{7, 14}))    // False
}
