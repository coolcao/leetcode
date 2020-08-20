/*
 * @lc app=leetcode id=1023 lang=golang
 *
 * [1023] Camelcase Matching
 *
 * https://leetcode.com/problems/camelcase-matching/description/
 *
 * algorithms
 * Medium (55.80%)
 * Likes:    212
 * Dislikes: 134
 * Total Accepted:    18.1K
 * Total Submissions: 31.7K
 * Testcase Example:  '["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]\n"FB"'
 *
 * A query word matches a given pattern if we can insert lowercase letters to
 * the pattern word so that it equals the query. (We may insert each character
 * at any position, and may insert 0 characters.)
 *
 * Given a list of queries, and a pattern, return an answer list of booleans,
 * where answer[i] is true if and only if queries[i] matches the pattern.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FB"
 * Output: [true,false,true,true,false]
 * Explanation:
 * "FooBar" can be generated like this "F" + "oo" + "B" + "ar".
 * "FootBall" can be generated like this "F" + "oot" + "B" + "all".
 * "FrameBuffer" can be generated like this "F" + "rame" + "B" + "uffer".
 *
 * Example 2:
 *
 *
 * Input: queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FoBa"
 * Output: [true,false,true,false,false]
 * Explanation:
 * "FooBar" can be generated like this "Fo" + "o" + "Ba" + "r".
 * "FootBall" can be generated like this "Fo" + "ot" + "Ba" + "ll".
 *
 *
 * Example 3:
 *
 *
 * Input: queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FoBaT"
 * Output: [false,true,false,false,false]
 * Explanation:
 * "FooBarTest" can be generated like this "Fo" + "o" + "Ba" + "r" + "T" +
 * "est".
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= queries.length <= 100
 * 1 <= queries[i].length <= 100
 * 1 <= pattern.length <= 100
 * All strings consists only of lower and upper case English letters.
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FoBaT"
	res := camelMatch(queries, pattern)
	fmt.Printf("%v\n", res)
}

// @lc code=start
func isLowerCase(c byte) bool {
	return c >= 'a' && c <= 'z'
}
func isUpperCase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}
func check(str string, pattern string) bool {
	if len(str) < len(pattern) {
		return false
	}
	is, ip := 0, 0
	for is < len(str) && ip < len(pattern) {
		if str[is] == pattern[ip] {
			is++
			ip++
			continue
		}
		if isLowerCase(str[is]) {
			is++
			continue
		}
		return false
	}
	if ip != len(pattern) {
		return false
	}
	for is < len(str) {
		if isUpperCase(str[is]) {
			return false
		}
		is++
	}
	return true
}
func camelMatch(queries []string, pattern string) []bool {
	ans := []bool{}
	for _, q := range queries {
		ans = append(ans, check(q, pattern))
	}
	return ans
}

// @lc code=end
