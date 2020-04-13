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
 *
 *
 *
 * Input: 16
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 14
 * Output: false
 *
 *
 *
 */
package main

import "fmt"

func main() {
	num := 334084
	fmt.Println(isPerfectSquare(num))
}

/*
问题：判断一个正整数是否是平方数。

判断一个数是否是平方数，最简单的办法就是从1开始计算平方，如果正好相等，那就是平方数。
但有一点需要注意，计算机中的整数，是有最大限制的，超过最大数会出现溢出。
所以在循环判断时，要注意计算平方数的溢出情况。
**/

// @lc code=start
func isPerfectSquare(num int) bool {
	// int 最大值
	INT_MAX := int(^uint(0) >> 1)
	i := 1
	isquare := i * i
	for isquare <= num && INT_MAX-isquare >= 2*i+1 {
		if isquare == num {
			return true
		}
		i++
		isquare = i * i
	}
	return false
}

// @lc code=end
