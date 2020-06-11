/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 *
 * https://leetcode.com/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (47.94%)
 * Likes:    1397
 * Dislikes: 198
 * Total Accepted:    203.2K
 * Total Submissions: 413.8K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * Given a string s and a string t, check if s is subsequence of t.
 *
 * A subsequence of a string is a new string which is formed from the original
 * string by deleting some (can be none) of the characters without disturbing
 * the relative positions of the remaining characters. (ie, "ace" is a
 * subsequence of "abcde" while "aec" is not).
 *
 * Follow up:
 * If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you
 * want to check one by one to see if T has its subsequence. In this scenario,
 * how would you change your code?
 *
 * Credits:
 * Special thanks to @pbrother for adding this problem and creating all test
 * cases.
 *
 *
 * Example 1:
 * Input: s = "abc", t = "ahbgdc"
 * Output: true
 *
 * Example 2:
 * Input: s = "axc", t = "ahbgdc"
 * Output: false
 *
 *
 * Constraints:
 *
 * 0 <= s.length <= 100
 * 0 <= t.length <= 10^4
 * Both strings consists only of lowercase characters.
 *
 *
 */
package main

import "fmt"

// @lc code=start
func isSubsequence(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls == 0 {
		return true
	}
	if lt == 0 {
		return false
	}
	si := 0
	for i := 0; i < lt && si < ls; i++ {
		if s[si] == t[i] {
			si++
		}
	}
	if si == ls {
		return true
	}
	return false
}

// @lc code=end

func main() {
	s := "acdf"
	t := "abcdef"
	is := isSubsequence(s, t)
	fmt.Printf("%v\n", is)
}
