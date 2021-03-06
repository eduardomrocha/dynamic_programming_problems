// How sum.

// Implementation of a solution to the how sum problem using memoization.

// Given an integer targetSum and a list of numbers as arguments, return a list containing
// any combination of elements that add up to exactly the targetSum. If there is no
// combination that adds up to the targetSum, then return None.
package main

import "fmt"

// Find a combination of elements from numbers that add up to exactly the targetSum.
func HowSum(targetSum int, numbers []int, memo *map[int][]int) []int {
	if combination, ok := (*memo)[targetSum]; ok {
		return combination
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, number := range numbers {
		remainder := targetSum - number
		combination := HowSum(remainder, numbers, memo)
		if combination != nil {
			combination = append(combination, number)
			(*memo)[targetSum] = combination
			return combination
		}
	}

	(*memo)[targetSum] = nil
	return nil
}

func main() {
	// Must create a map for each run of HowSum, because golang has no default value for
	// function params or optional params in function. If we pass the same map as param
	// to each run of HowSum, the values stored in the first run will affect the
	// subsequent ones because not all HowSum calls have the same targetSum. Another
	// solution would be to create just one map and after each run delete all results
	// stored inside it.
	memo1 := make(map[int][]int)
	memo2 := make(map[int][]int)
	memo3 := make(map[int][]int)
	memo4 := make(map[int][]int)
	memo5 := make(map[int][]int)

	fmt.Println(HowSum(7, []int{2, 3}, &memo1))       // [3,2,2]
	fmt.Println(HowSum(7, []int{5, 3, 4, 7}, &memo2)) // [4,3]
	fmt.Println(HowSum(7, []int{2, 4}, &memo3))       // []
	fmt.Println(HowSum(8, []int{2, 3, 5}, &memo4))    // [2,2,2,2]
	fmt.Println(HowSum(300, []int{7, 14}, &memo5))    // []
}
