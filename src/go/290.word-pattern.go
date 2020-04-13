/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 *
 * https://leetcode.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (36.45%)
 * Likes:    940
 * Dislikes: 131
 * Total Accepted:    177.1K
 * Total Submissions: 485.8K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 *
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in str.
 *
 * Example 1:
 *
 *
 * Input: pattern = "abba", str = "dog cat cat dog"
 * Output: true
 *
 * Example 2:
 *
 *
 * Input:pattern = "abba", str = "dog cat cat fish"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: pattern = "aaaa", str = "dog cat cat dog"
 * Output: false
 *
 * Example 4:
 *
 *
 * Input: pattern = "abba", str = "dog dog dog dog"
 * Output: false
 *
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains
 * lowercase letters that may be separated by a single space.
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abcad"
	str := "dog cat ca dog cate"
	fmt.Println(wordPattern(pattern, str))
}

// @lc code=start
func wordPattern(pattern string, str string) bool {
	patterns := strings.Split(pattern, "")
	words := strings.Split(str, " ")
	pl, wl := len(patterns), len(words)
	if pl != wl {
		return false
	}

	m1 := map[string]string{}
	m2 := map[string]string{}

	for i := 0; i < pl; i++ {
		p, w := patterns[i], words[i]
		_w, ok1 := m1[p]
		_p, ok2 := m2[w]
		if !ok1 && !ok2 {
			m1[p] = w
			m2[w] = p
		} else if ok1 && ok2 {
			if p == _p && w == _w {
				continue
			}
			return false
		} else {
			return false
		}
	}

	return true
}

// @lc code=end
