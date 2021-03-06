// Grid Traveler.

// Implementation of a solution to the grid traveler problem using memoization.

// Say that you are a traveler on a 2D grid. You begin in the top-left corner and
// your goal is to travel to the bottom-right corner. You may only move down or right.

// In how many ways can you travel to the goal on a grid with dimensions m * n?
package main

import (
	"fmt"
	"strconv"
)

// Given the number of lines and columns, return the number of ways the traveler can
// travel to the goal.
func GridTraveler(m int, n int, memo *map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)

	if ans, ok := (*memo)[key]; ok {
		return ans
	}

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	(*memo)[key] = GridTraveler(m-1, n, memo) + GridTraveler(m, n-1, memo)
	return (*memo)[key]

}

func main() {
	memo := make(map[string]int)

	fmt.Println(GridTraveler(1, 1, &memo))   // 1
	fmt.Println(GridTraveler(2, 3, &memo))   // 3
	fmt.Println(GridTraveler(3, 2, &memo))   // 3
	fmt.Println(GridTraveler(3, 3, &memo))   // 6
	fmt.Println(GridTraveler(18, 18, &memo)) // 2333606220
}
