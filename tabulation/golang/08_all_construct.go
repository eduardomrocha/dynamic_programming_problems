// All Construct.

// Implementation of a solution to the all construct problem using tabulation.

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
func AllConstruct(target string, wordBank []string) [][]string {
	var table [][][]string

	for i := 0; i <= len(target); i++ {
		row := [][]string{nil}
		table = append(table, row)
	}

	table[0][0] = []string{}

	for i := 0; i <= len(target); i++ {
		if table[i][0] != nil {
			for _, word := range wordBank {
				if strings.HasPrefix(target[i:], word) {
					for _, combination := range table[i] {
						combination = append(combination, word)
						if table[i+len(word)][0] == nil {
							table[i+len(word)][0] = combination
						} else {
							table[i+len(word)] = append(table[i+len(word)], combination)
						}
					}
				}
			}
		}
	}

	return table[len(target)]
}

func main() {
	fmt.Println(AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	/*
	   [
	       ["purp", "le"],
	       ["p", "ur", "p", "le"]
	   ]
	*/
	fmt.Println(AllConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	/*
	   [
	       ["ab", "cd", "ef"],
	       ["ab", "c", "def"],
	       ["abc", "def"],
	       ["abcd", "ef"]
	   ]
	*/
	fmt.Println(AllConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                               // []
	fmt.Println(AllConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // []
}
