/*
 * @lc app=leetcode id=1137 lang=golang
 *
 * [1137] N-th Tribonacci Number
 *
 * https://leetcode.com/problems/n-th-tribonacci-number/description/
 *
 * algorithms
 * Easy (56.74%)
 * Likes:    298
 * Dislikes: 41
 * Total Accepted:    42.7K
 * Total Submissions: 76.3K
 * Testcase Example:  '4'
 *
 * The Tribonacci sequence Tn is defined as follows:
 *
 * T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
 *
 * Given n, return the value of Tn.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4
 * Output: 4
 * Explanation:
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 *
 *
 * Example 2:
 *
 *
 * Input: n = 25
 * Output: 1389537
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 37
 * The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31
 * - 1.
 *
 */
package main

import "fmt"

func main() {
	n := 0
	fmt.Printf("%v\n", tribonacci(n))
}

// @lc code=start
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0], nums[1], nums[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2] + nums[i-3]
	}
	return nums[n]
}

// @lc code=end
