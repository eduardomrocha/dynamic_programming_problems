// Grid traveler.

// Implementation of a solution to the grid traveler problem using tabulation.

// Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal
// is to travel to the bottom-right corner. You may only move down or right.

// In how many ways can you travel to the goal on a grid with dimensions m * n?
package main

import "fmt"

// Given the number of lines and columns, return the number of ways the traveler can
// travel to the goal.
func GridTraveler(m int, n int) int {
	var table [][]int

	for i := 0; i <= m; i++ {
		var row []int
		for j := 0; j <= n; j++ {
			row = append(row, 0)
		}
		table = append(table, row)
	}

	table[1][1] = 1

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j+1 <= n {
				table[i][j+1] += table[i][j]
			}

			if i+1 <= m {
				table[i+1][j] += table[i][j]
			}
		}
	}

	return table[m][n]
}

func main() {
	fmt.Println(GridTraveler(1, 1))   // 1
	fmt.Println(GridTraveler(2, 3))   // 3
	fmt.Println(GridTraveler(3, 2))   // 3
	fmt.Println(GridTraveler(3, 3))   // 6
	fmt.Println(GridTraveler(18, 18)) // 2333606220
}
