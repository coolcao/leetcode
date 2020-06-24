/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 *
 * https://leetcode.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (35.95%)
 * Likes:    1570
 * Dislikes: 100
 * Total Accepted:    163.8K
 * Total Submissions: 449.9K
 * Testcase Example:  '"aba"'
 *
 *
 * Given a non-empty string s, you may delete at most one character.  Judge
 * whether you can make it a palindrome.
 *
 *
 * Example 1:
 *
 * Input: "aba"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 *
 *
 *
 * Note:
 *
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 *
 *
 */
package main

import "fmt"

func main() {
	s := "abab"
	valid := validPalindrome(s)
	fmt.Printf("%v\n", valid)
}

// @lc code=start
func check(s string, start, end int) bool {
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
func validPalindrome(s string) bool {
	length := len(s)
	start, end := 0, length-1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return check(s, start+1, end) || check(s, start, end-1)
		}

	}
	return true
}

// @lc code=end
