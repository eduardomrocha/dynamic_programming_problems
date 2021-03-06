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
func CanConstruct(target string, wordBank []string) bool {
	var table []bool

	for i := 0; i <= len(target); i++ {
		table = append(table, false)
	}

	table[0] = true

	for i := 0; i <= len(target); i++ {
		if table[i] {
			for _, word := range wordBank {
				if strings.HasPrefix(target[i:], word) {
					table[i+len(word)] = true
				}
			}
		}
	}

	return table[len(target)]
}

func main() {
	fmt.Println(CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                                              // true
	fmt.Println(CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                               // false
	fmt.Println(CanConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))                             // true
	fmt.Println(CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // false
}
