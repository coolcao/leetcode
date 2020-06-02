/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (33.66%)
 * Likes:    3938
 * Dislikes: 276
 * Total Accepted:    365.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 *
 * Example:
 *
 *
 * Input: S = "ADOBECODEBBANC", T = "ABC"
 * Output: "BANC"
 *
 *
 * Note:
 *
 *
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 *
 *
 */
package main

import "fmt"

func main() {
	s := "AAAAA"
	t := "AA"
	min := minWindow(s, t)
	fmt.Printf("%v\n", min)
}

// @lc code=start
func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if ls < lt {
		return ""
	}

	// 窗口里存的是t中字符应该出现的次数
	// 正数表示该字符还缺的出现次数，0表示刚好出现，负数表示s中字符出现的次数多于t中字符出现次数
	windowMap := map[byte]int{}

	// 初始化窗口
	for i := 0; i < lt; i++ {
		windowMap[t[i]]++
	}

	windowSize := len(windowMap)
	left, right := 0, 0

	// 窗口中已经包含T的不同字符的种类
	c := 0
	ans := ""

	for right < ls {
		windowMap[s[right]]--
		// 统计窗口中已经包含的T中的不同字符的种类
		if windowMap[s[right]] == 0 {
			c++
		}
		for c == windowSize && windowMap[s[left]] < 0 {
			windowMap[s[left]]++
			left++
		}
		if c == windowSize {
			if len(ans) == 0 || right-left+1 < len(ans) {
				ans = s[left : right+1]
			}
		}
		right++
	}

	return ans
}

// @lc code=end
