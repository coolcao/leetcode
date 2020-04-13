/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 *
 * https://leetcode.com/problems/sum-of-two-integers/description/
 *
 * algorithms
 * Easy (50.65%)
 * Likes:    1084
 * Dislikes: 1987
 * Total Accepted:    173.3K
 * Total Submissions: 342.3K
 * Testcase Example:  '1\n2'
 *
 * Calculate the sum of two integers a and b,
 * but you are not allowed to use,the operator + and -.
 * 计算整数a和b的和，但是不允许使用 + 和 -
 *
 * Example 1:
 * Input: a = 1, b = 2
 * Output: 3
 *
 * Example 2:
 * Input: a = -2, b = 3
 * Output: 1
 *
 */
package main

import "fmt"

func main() {
	a, b := 32, 654
	fmt.Printf("%d+%d=%d\n", a, b, getSum(a, b))
}

// @lc code=start
func getSum(a int, b int) int {
	// 当进位部分为0时，即返回
	if b == 0 {
		return a
	}
	sum := a ^ b        // 不进位部分
	carry := a & b << 1 // 进位部分
	return getSum(sum, carry)
}

// @lc code=end
