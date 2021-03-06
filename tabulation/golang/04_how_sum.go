// How sum.

// Implementation of a solution to the how sum problem using tabulation.

// Given an integer targetSum and a list of numbers as arguments, return a list containing
// any combination of elements that add up to exactly the targetSum. If there is no
// combination that adds up to the targetSum, then return None.
package main

import "fmt"

// Find a combination of elements from numbers that add up to exactly the targetSum.
func HowSum(targetSum int, numbers []int) []int {
	var table [][]int

	for i := 0; i <= targetSum; i++ {
		table = append(table, nil)
	}

	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, number := range numbers {
				if i+number <= targetSum {
					table[i+number] = append(table[i], number)
				}
			}
		}
	}

	return table[targetSum]
}

func main() {
	fmt.Println(HowSum(7, []int{2, 3}))       // [3,2,2]
	fmt.Println(HowSum(7, []int{5, 3, 4, 7})) // [4,3]
	fmt.Println(HowSum(7, []int{2, 4}))       // []
	fmt.Println(HowSum(8, []int{2, 3, 5}))    // [2,2,2,2]
	fmt.Println(HowSum(300, []int{7, 14}))    // []
}
