// Can construct.

// Implementation of a solution to the can construct problem using memoization.

// Given a string 'target' and a list of strings 'wordBank', return a boolean indicating
// whether or not the 'target' can be constructed by concatenating elements of 'wordBank'
// list.

// You may reuse elements of 'wordBank' as many times as needed.
package main

import (
	"fmt"
	"strings"
)

// Check if a given target can be constructed with a given word bank.
func CanConstruct(target string, wordBank []string, memo *map[string]bool) bool {
	if ans, ok := (*memo)[target]; ok {
		return ans
	}

	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			if CanConstruct(target[len(word):], wordBank, memo) {
				(*memo)[target] = true
				return true
			}
		}
	}

	(*memo)[target] = false
	return false
}

func main() {
	// Must create a map for each run of CanConstruct, because golang has no default value for
	// function params or optional params in function. If we pass the same map as param
	// to each run of CanConstruct, the values stored in the first run will affect the
	// subsequent ones because not all CanConstruct calls have the same targetSum. Another
	// solution would be to create just one map and after each run delete all results
	// stored inside it.
	memo1 := make(map[string]bool)
	memo2 := make(map[string]bool)
	memo3 := make(map[string]bool)
	memo4 := make(map[string]bool)

	fmt.Println(CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, &memo1))                                              // true
	fmt.Println(CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, &memo2))                               // false
	fmt.Println(CanConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, &memo3))                             // true
	fmt.Println(CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, &memo4)) // false
}
