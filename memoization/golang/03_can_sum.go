// Can sum.

// Implementation of a solution to the can sum problem using memoization.

// Given an integer targetSum and a list of numbers as arguments, return a boolean
// indicating whether or not it is possible to generate the targetSum using numbers from
// the list, using an element of the list as many times as needed and assuming that all
// input numbers are nonnegative.
package main

import "fmt"

// Check if the targetSum can be gereteded by adding up elements from numbers.
func CanSum(targetSum int, numbers []int, memo *map[int]bool) bool {
	if ans, ok := (*memo)[targetSum]; ok {
		return ans
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, number := range numbers {
		remainder := targetSum - number
		if CanSum(remainder, numbers, memo) {
			(*memo)[targetSum] = true
			return true
		}
	}

	(*memo)[targetSum] = false
	return false
}

func main() {
	// Must create a map for each run of CanSum, because golang has no default value for
	// function params or optional params in function. If we pass the same map as param
	// to each run of CanSum, the values stored in the first run will affect the
	// subsequent ones because not all CanSum calls have the same targetSum. Another
	// solution would be to create just one map and after each run delete all results
	// stored inside it.
	memo1 := make(map[int]bool)
	memo2 := make(map[int]bool)
	memo3 := make(map[int]bool)
	memo4 := make(map[int]bool)
	memo5 := make(map[int]bool)

	fmt.Println(CanSum(7, []int{2, 3}, &memo1))       // True
	fmt.Println(CanSum(7, []int{5, 3, 4, 7}, &memo2)) // True
	fmt.Println(CanSum(7, []int{2, 4}, &memo3))       // False
	fmt.Println(CanSum(8, []int{2, 3, 5}, &memo4))    // True
	fmt.Println(CanSum(300, []int{7, 14}, &memo5))    // False
}
