package main

import "fmt"

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/

func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	start := 0
	result := make([]int, len(s))
	i := 0
	for idx, val := range s {
		if val == ' ' {
			result[i] = idx - start
			start = idx + 1
			i = i + 1
		} else if idx == len(s)-1 {
			result[i] = idx - start + 1
			start = idx + 1
			i = i + 1
		}
	}
	fmt.Printf("%v\n", result)
	for i > 0 {
		i--
		if result[i] > 0 {
			return result[i]
		}
	}
	return 0
}

func main() {
	s := " "
	length := lengthOfLastWord(s)
	fmt.Printf("%d\n", length)
}
