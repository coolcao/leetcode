/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...
Example 1:

Input: 1
Output: "A"
Example 2:

Input: 28
Output: "AB"
Example 3:

Input: 701
Output: "ZY"
*/

package main

import "fmt"

func getCharIndex(n int) []int {
	idx := make([]int, 0)
	if n <= 26 {
		idx = append(idx, n)
		return idx
	}
	for ; n > 26; n = (int)(n / 26) {
		mod := n % 26
		if mod == 0 {
			idx = append(idx, 26)
			n--
		} else {
			idx = append(idx, mod)
		}
	}
	idx = append(idx, n)
	return idx
}

func convertToTitle(n int) string {
	chars := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	idx := getCharIndex(n)
	length := len(idx)
	result := ""
	for i := length - 1; i >= 0; i-- {
		result += chars[idx[i]]
	}
	return result
}

func main() {
	// n := 676
	// result := convertToTitle(n)
	// fmt.Printf("%s\n", result)
	for i := 1; i < 10000; i++ {
		fmt.Printf("i:%d\t%s\n", i, convertToTitle(i))
	}
}

// @lc code=end
