/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (40.97%)
 * Likes:    647
 * Dislikes: 137
 * Total Accepted:    150.3K
 * Total Submissions: 366.7K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 *
 * Note: Do not use any built-in library function such as sqrt.
 *
 * Example 1:
 * Input: 16
 * Output: true
 *
 *
 *
 * Example 2:
 * Input: 14
 * Output: false
 *
 */
package main

// @lc code=start
func isPerfectSquare(num int) bool {
	// 连续奇数的和即为完全平方数
	for i := 1; num > 0; i += 2 {
		num = num - i
	}
	return num == 0
}

// @lc code=end
