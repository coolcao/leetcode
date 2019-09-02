package main

import "fmt"

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	minLen := len(strs[0])
	strLen := len(strs)
	tmpMap := make(map[int]string)

	for idx, s := range strs {
		length := len(s)
		if minLen > length {
			minLen = length
		}
		tmpMap[idx] = s
	}
	for i := 0; i < minLen; i++ {
		firstChar := tmpMap[0][i]
		for j := 1; j < strLen; j++ {
			str := tmpMap[j]
			if firstChar == str[i] {
				continue
			} else {
				return result
			}
		}
		result += string(firstChar)
	}
	return result
}

func main() {
	strs := []string{"flower", "flow", "floight"}
	result := longestCommonPrefix(strs)
	fmt.Printf("result = %s\n", result)
}
