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
func CountConstruct(target string, wordBank []string) int {
	var table []int

	for i := 0; i <= len(target); i++ {
		table = append(table, 0)
	}

	table[0] = 1

	for i := 0; i <= len(target); i++ {
		if table[i] > 0 {
			for _, word := range wordBank {
				if strings.HasPrefix(target[i:], word) {
					if i+len(word) <= len(target) {
						table[i+len(word)] += table[i]
					}
				}
			}
		}
	}

	return table[len(target)]
}

func main() {
	fmt.Println(CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                                              // 2
	fmt.Println(CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                                              // 1
	fmt.Println(CountConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                               // 0
	fmt.Println(CountConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))                             // 4
	fmt.Println(CountConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // 0
}
