// All Construct.

// Implementation of a solution to the all construct problem using memoization.

// Given a string 'target' and a list of strings 'wordBank', return a 2D list containing
// all of the ways that 'target' can be constructed by concatenating elements of the
// 'wordBank' list. Each element of the 2D list should represent one combination that
// constructs the 'target'.

// You may reuse elements of 'wordBank' as many times as needed.
package main

import (
	"fmt"
	"strings"
)

// Construct all ways that target can be constructed using the words in wordBank.
func AllConstruct(target string, wordBank []string, memo *map[string][][]string) [][]string {
	if ans, ok := (*memo)[target]; ok {
		return ans
	}

	if target == "" {
		var ans [][]string
		ans = append(ans, []string{})

		return ans
	}

	var result [][]string

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffixWays := AllConstruct(target[len(word):], wordBank, memo)
			var targetWays [][]string

			for _, way := range suffixWays {
				newWay := append([]string{word}, way...)
				targetWays = append(targetWays, newWay)
			}

			result = append(result, targetWays...)
		}
	}

	(*memo)[target] = result
	return result
}

func main() {
	// Must create a map for each run of AllConstruct, because golang has no default value for
	// function params or optional params in function. If we pass the same map as param
	// to each run of AllConstruct, the values stored in the first run will affect the
	// subsequent ones because not all AllConstruct calls have the same targetSum. Another
	// solution would be to create just one map and after each run delete all results
	// stored inside it.
	memo1 := make(map[string][][]string)
	memo2 := make(map[string][][]string)
	memo3 := make(map[string][][]string)
	memo4 := make(map[string][][]string)

	fmt.Println(AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, &memo1))
	/*
	   [
	       ["purp", "le"],
	       ["p", "ur", "p", "le"]
	   ]
	*/
	fmt.Println(AllConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, &memo2))
	/*
	   [
	       ["ab", "cd", "ef"],
	       ["ab", "c", "def"],
	       ["abc", "def"],
	       ["abcd", "ef"]
	   ]
	*/
	fmt.Println(AllConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, &memo3))                               // []
	fmt.Println(AllConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, &memo4)) // []
}
