// Best sum.

// Implementation of a solution to the best sum problem using tabulation.

// Given an integer targetSum and a list of numbers as arguments, return a list containing
// the shortest combination of elements that add up to exactly the targetSum. If there is
// a tie for the shortest combination, you may return any one of the shortest.
package main

import "fmt"

// Find the best combination of elements that add up to exactly the targetSum.
func BestSum(targetSum int, numbers []int) []int {
	var table [][]int

	for i := 0; i <= targetSum; i++ {
		table = append(table, nil)
	}

	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, number := range numbers {
				if i+number <= targetSum {
					newCombination := append(table[i], number)
					if table[i+number] == nil || len(table[i+number]) > len(newCombination) {
						table[i+number] = newCombination
					}
				}
			}
		}
	}

	return table[targetSum]
}

func main() {
	fmt.Println(BestSum(7, []int{5, 3, 4, 7}))    // [7]
	fmt.Println(BestSum(8, []int{2, 3, 5}))       // [3, 5]
	fmt.Println(BestSum(8, []int{1, 4, 5}))       // [4, 4]
	fmt.Println(BestSum(100, []int{1, 3, 5, 25})) // [25, 25, 25, 25]
}
