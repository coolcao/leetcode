/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 *
 * https://leetcode.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (40.37%)
 * Likes:    1128
 * Dislikes: 52
 * Total Accepted:    84.9K
 * Total Submissions: 209.4K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * Given two strings s1 and s2, write a function to return true if s2 contains
 * the permutation of s1. In other words, one of the first string's
 * permutations is the substring of the second string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "ab" s2 = "eidbaooo"
 * Output: True
 * Explanation: s2 contains one permutation of s1 ("ba").
 *
 *
 * Example 2:
 *
 *
 * Input:s1= "ab" s2 = "eidboaoo"
 * Output: False
 *
 *
 *
 *
 * Note:
 *
 *
 * The input strings only contain lower case letters.
 * The length of both given strings is in range [1, 10,000].
 *
 *
 */
package main

import "fmt"

func main() {
	s1 := "abfc"
	s2 := "eidcfbaagooo"
	exist := checkInclusion(s1, s2)
	fmt.Println(exist)
}

// @lc code=start

/*
用哈希表记录窗口中各个字符出现次数的差值
正数表示还应该出现几次
0表示正好
负数表示多出现了几次
**/
func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}

	windowMap := map[byte]int{}

	// 初始化窗口
	for i := 0; i < l1; i++ {
		windowMap[s1[i]]++
	}

	left, right := 0, 0
	for right < l2 {
		c := s2[right]
		// 入窗
		windowMap[c]--
		right++
		// 出窗。
		for left < right && windowMap[c] < 0 {
			windowMap[s2[left]]++
			left++
		}
		if right-left == l1 {
			return true
		}
	}

	return false
}

// @lc code=end
