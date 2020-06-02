/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.77%)
 * Likes:    8970
 * Dislikes: 541
 * Total Accepted:    1.5M
 * Total Submissions: 5M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 * Example 2:
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 * Example 3:
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 */
package main

import "fmt"

func main() {
	s := "abcabcdd"
	max := lengthOfLongestSubstring(s)
	fmt.Printf("%v\n", max)
}

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	length := len(chars)
	if length == 0 {
		return 0
	}

	max := 0
	hmap := map[rune]int{}

	left, right := 0, 0
	for right < length {
		c := chars[right]
		// 入窗
		hmap[c]++
		right++

		for left < right && hmap[c] > 1 {
			hmap[chars[left]]--
			left++
		}
		if max < right-left {
			max = right - left
		}

	}
	return max
}

// @lc code=end
