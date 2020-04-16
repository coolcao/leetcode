/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 *
 * https://leetcode.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (41.68%)
 * Likes:    1257
 * Dislikes: 137
 * Total Accepted:    113.8K
 * Total Submissions: 272.9K
 * Testcase Example:  '"abab"'
 *
 * Given a non-empty string check if it can be constructed by taking a
 * substring of it and appending multiple copies of the substring together. You
 * may assume the given string consists of lowercase English letters only and
 * its length will not exceed 10000.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "abab"
 * Output: True
 * Explanation: It's the substring "ab" twice.
 *
 *
 * Example 2:
 *
 *
 * Input: "aba"
 * Output: False
 *
 *
 * Example 3:
 *
 *
 * Input: "abcabcabcabc"
 * Output: True
 * Explanation: It's the substring "abc" four times. (And the substring
 * "abcabc" twice.)
 *
 *
 */
package main

import "fmt"

func main() {
	s := "abcdefgabcdefghh"
	b := repeatedSubstringPattern(s)
	fmt.Println(b)
}

// @lc code=start
// 方法1:设定一个步长step，依次检查这个步长的子串是否重复。
// 步长从1到length/2，因为步长为length/2时，正好是两个子串的重复。step不可能再大于length/2了
func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for step := 1; step <= length/2; step++ {
		if length%step != 0 {
			continue
		}
	L2:
		for i := 0; i < step; i++ {
			char := s[i]
			for j := i; j < length; j += step {
				if i == step-1 && j == length-step+i {
					if char == s[j] {
						return true
					}
					return false
				}
				if char == s[j] {
					continue
				} else {
					break L2
				}
			}
		}
	}
	return false
}

// 方法2:利用map优化上面的方法
// 上面方法中，step从1开始，依次到length/2
// 假设step=8时，字符串不符合规范，那么8的因子1,2,4将都不符合。因为1重复8次，2重复4次，4重复2次即可得到8
// 因此，我们可以将step从大到小检查。如果不符合规范，放到一个map里
// 检查一个step是否符合规范时，先检查这个map里的step是否是step的倍数，如果是，那么这个step肯定是不符合规范的。
func repeatedSubstringPattern2(s string) bool {
	length := len(s)
	m := map[int]int{}
	for step := length / 2; step >= 1; step-- {
		if length%step != 0 {
			continue
		}
	L1:
		for _, v := range m {
			if v%step == 0 {
				break L1
			}
		}
	L2:
		for i := 0; i < step; i++ {
			char := s[i]
			for j := i; j < length; j += step {
				if i == step-1 && j == length-step+i {
					if char == s[j] {
						return true
					}
					return false
				}
				if char == s[j] {
					continue
				} else {
					m[step] = step
					break L2
				}
			}
		}
	}
	return false
}

// @lc code=end
