/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 *
 * https://leetcode.com/problems/rotate-string/description/
 *
 * algorithms
 * Easy (49.55%)
 * Likes:    680
 * Dislikes: 50
 * Total Accepted:    65.4K
 * Total Submissions: 131.5K
 * Testcase Example:  '"abcde"\n"cdeab"'
 *
 * We are given two strings, A and B.
 *
 * A shift on A consists of taking string A and moving the leftmost character
 * to the rightmost position. For example, if A = 'abcde', then it will be
 * 'bcdea' after one shift on A. Return True if and only if A can become B
 * after some number of shifts on A.
 *
 *
 * Example 1:
 * Input: A = 'abcde', B = 'cdeab'
 * Output: true
 *
 * Example 2:
 * Input: A = 'abcde', B = 'abced'
 * Output: false
 *
 *
 * Note:
 *
 *
 * A and B will have length at most 100.
 *
 *
 */
package main

import (
	"fmt"
	"strings"
)

func rotateString2(A string, B string) bool {
	la, lb := len(A), len(B)
	if la != lb {
		return false
	}
	start := 0
	// 先找到旋转的点
	for i := 0; i < la; i++ {
		if A[i] == B[0] {
			j := i
			for j < la {
				if A[j] != B[j-i] {
					break
				}
				j++
			}
			if j == la {
				start = i
				break
			}
		}
	}
	// 分别判断旋转的点的前后是否相等即可
	for i := start; i < la; i++ {
		if A[i] != B[i-start] {
			return false
		}
	}
	for i := 0; i < start; i++ {
		if A[i] != B[lb-start+i] {
			return false
		}
	}
	return true
}

// @lc code=start
func rotateString(A string, B string) bool {
	la, lb := len(A), len(B)
	if la != lb {
		return false
	}
	if la == 0 {
		return true
	}

	AA := A + A

	if strings.Contains(AA, B) {
		return true
	}

	return false
}

// @lc code=end

func main() {
	A, B := "abcde", "abced"
	b := rotateString(A, B)
	fmt.Printf("%v\n", b)
}
