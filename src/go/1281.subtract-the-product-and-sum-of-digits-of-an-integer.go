/*
 * @lc app=leetcode id=1281 lang=golang
 *
 * [1281] Subtract the Product and Sum of Digits of an Integer
 *
 * https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/description/
 *
 * algorithms
 * Easy (84.65%)
 * Likes:    270
 * Dislikes: 94
 * Total Accepted:    79.1K
 * Total Submissions: 93.2K
 * Testcase Example:  '234'
 *
 * Given an integer number n, return the difference between the product of its
 * digits and the sum of its digits.
 *
 * Example 1:
 *
 *
 * Input: n = 234
 * Output: 15
 * Explanation:
 * Product of digits = 2 * 3 * 4 = 24
 * Sum of digits = 2 + 3 + 4 = 9
 * Result = 24 - 9 = 15
 *
 *
 * Example 2:
 *
 *
 * Input: n = 4421
 * Output: 21
 * Explanation:
 * Product of digits = 4 * 4 * 2 * 1 = 32
 * Sum of digits = 4 + 4 + 2 + 1 = 11
 * Result = 32 - 11 = 21
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^5
 *
 *
 */
package main

import "fmt"

func main() {
	n := 11
	fmt.Printf("%v\n", subtractProductAndSum(n))
}

// @lc code=start
func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n > 0 {
		mod := n % 10
		product = product * mod
		sum = sum + mod
		n = n / 10
	}
	return product - sum
}

// @lc code=end
