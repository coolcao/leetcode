package main

import "fmt"

/*
28. Implement strStr()

Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().


*/

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	hayLen, needLen := len(haystack), len(needle)

	idx := -1
OUTER:
	for i := 0; i <= hayLen-needLen; i++ {
		idx = -1
		if haystack[i] != needle[0] {
			continue OUTER
		}
		idx = i
	INNER:
		for j := 0; j < needLen; j++ {
			if haystack[i+j] == needle[j] {
				if j == needLen-1 {
					return idx
				}
				continue INNER
			}
			if i == hayLen-needLen {
				idx = -1
			}
			continue OUTER
		}
	}
	return idx
}

func main() {
	haystack := "mississippi"
	needle := "issipi"
	result := strStr(haystack, needle)
	fmt.Printf("%d\n", result)
}
