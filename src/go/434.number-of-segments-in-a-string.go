/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 *
 * https://leetcode.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (37.49%)
 * Likes:    189
 * Dislikes: 686
 * Total Accepted:    71.3K
 * Total Submissions: 190.1K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * Count the number of segments in a string, where a segment is defined to be a
 * contiguous sequence of non-space characters.
 *
 * Please note that the string does not contain any non-printable characters.
 *
 * Example:
 *
 * Input: "Hello, my name is John"
 * Output: 5
 *
 *
 */
package main

import "fmt"

// import (
// 	"fmt"
// 	"strings"
// )

func main() {
	s := "a, b, c"
	fmt.Printf("count: %d\n", countSegments(s))
}

// @lc code=start
// func countSegments(s string) int {
// 	ss := strings.Split(s, " ")
// 	count := 0
// 	for i := 0; i < len(ss); i++ {
// 		if ss[i] != "" {
// 			count++
// 		}
// 	}
// 	return count
// }

func countSegments(s string) int {
	chars := []rune(s)
	length := len(chars)
	if length == 0 {
		return 0
	}
	space := rune(' ')
	pre := chars[0]
	count := 0
	for i := 0; i < length; i++ {
		c := chars[i]
		if i == length-1 && c != space {
			count++
		}
		if c == space && pre != space {
			count++
		}

		pre = c
	}
	return count
}

// @lc code=end
