// Count construct.

// Implementation of a solution to the count construct problem using memoization.

// Given a string 'target' and a list of strings 'wordBank', return the number of ways
// that 'target' can be constructed by concatenating elements of the 'wordBank' list.

// You may reuse elements of 'wordBank' as many times as needed.
package main

import (
	"fmt"
	"strings"
)

// Count in how many ways the target can be constructed with a given word bank.
func CountConstruct(target string, wordBank []string, memo *map[string]int) int {
	if ans, ok := (*memo)[target]; ok {
		return ans
	}

	if target == "" {
		return 1
	}

	count := 0

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffixCount := CountConstruct(target[len(word):], wordBank, memo)
			count += suffixCount
		}
	}

	(*memo)[target] = count
	return count
}

func main() {
	// Must create a map for each run of CountConstruct, because golang has no default value for
	// function params or optional params in function. If we pass the same map as param
	// to each run of CountConstruct, the values stored in the first run will affect the
	// subsequent ones because not all CountConstruct calls have the same targetSum. Another
	// solution would be to create just one map and after each run delete all results
	// stored inside it.
	memo1 := make(map[string]int)
	memo2 := make(map[string]int)
	memo3 := make(map[string]int)
	memo4 := make(map[string]int)
	memo5 := make(map[string]int)

	fmt.Println(CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, &memo1))                                              // 2
	fmt.Println(CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, &memo2))                                              // 1
	fmt.Println(CountConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, &memo3))                               // 0
	fmt.Println(CountConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, &memo4))                             // 4
	fmt.Println(CountConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, &memo5)) // 0
}
