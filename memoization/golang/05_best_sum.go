// Best sum.

// Implementation of a solution to the best sum problem using memoization.

// Given an integer targetSum and a list of numbers as arguments, return a list containing
// the shortest combination of elements that add up to exactly the targetSum. If there is
// a tie for the shortest combination, you may return any one of the shortest.
package main

import "fmt"

// Find the best combination of elements that add up to exactly the targetSum.
func BestSum(targetSum int, numbers []int, memo *map[int][]int) []int {
	if combination, ok := (*memo)[targetSum]; ok {
		return combination
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	var bestCombination []int

	for _, number := range numbers {
		remainder := targetSum - number
		combination := BestSum(remainder, numbers, memo)

		if combination != nil {
			combination = append(combination, number)

			if bestCombination == nil || len(combination) < len(bestCombination) {
				bestCombination = combination
			}
		}
	}

	(*memo)[targetSum] = bestCombination
	return bestCombination
}

func main() {
	// Must create a map for each run of BestSum, because golang has no default value for
	// function params or optional params in function. If we pass the same map as param
	// to each run of BestSum, the values stored in the first run will affect the
	// subsequent ones because not all BestSum calls have the same targetSum. Another
	// solution would be to create just one map and after each run delete all results
	// stored inside it.
	memo1 := make(map[int][]int)
	memo2 := make(map[int][]int)
	memo3 := make(map[int][]int)
	memo4 := make(map[int][]int)

	fmt.Println(BestSum(7, []int{5, 3, 4, 7}, &memo1))    // [7]
	fmt.Println(BestSum(8, []int{2, 3, 5}, &memo2))       // [3, 5]
	fmt.Println(BestSum(8, []int{1, 4, 5}, &memo3))       // [4, 4]
	fmt.Println(BestSum(100, []int{1, 3, 5, 25}, &memo4)) // [25, 25, 25, 25]
}
