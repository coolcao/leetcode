/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (54.98%)
 * Likes:    3813
 * Dislikes: 193
 * Total Accepted:    718.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 *
 * Example:
 *
 *
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * Note:
 *
 *
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Printf("%v\n", ans)
}

// @lc code=start
// 排序
func sort(str string) string {
	if str == "" {
		return ""
	}
	counts := [26]int{}
	for _, c := range str {
		idx := int(c) - 97
		counts[idx]++
	}
	result := []rune{}
	for idx, c := range counts {
		if c > 0 {
			for i := 0; i < c; i++ {
				result = append(result, rune(idx+97))
			}
		}
	}
	return string(result)
}

func groupAnagrams(strs []string) [][]string {
	hmap := map[string][]string{}

	for _, str := range strs {
		sorted := sort(str)
		if _, ok := hmap[sorted]; ok {
			hmap[sorted] = append(hmap[sorted], str)
		} else {
			hmap[sorted] = []string{str}
		}
	}

	ans := [][]string{}
	for _, tmp := range hmap {
		ans = append(ans, tmp)
	}

	return ans
}

// @lc code=end
