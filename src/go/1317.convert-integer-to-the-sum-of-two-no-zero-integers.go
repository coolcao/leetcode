/*
 * @lc app=leetcode id=1317 lang=golang
 *
 * [1317] Convert Integer to the Sum of Two No-Zero Integers
 *
 * https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/description/
 *
 * algorithms
 * Easy (56.91%)
 * Likes:    97
 * Dislikes: 100
 * Total Accepted:    14.8K
 * Total Submissions: 26.2K
 * Testcase Example:  '2'
 *
 * Given an integer n. No-Zero integer is a positive integer which doesn't
 * contain any 0 in its decimal representation.
 *
 * Return a list of two integers [A, B] where:
 *
 *
 * A and B are No-Zero integers.
 * A + B = n
 *
 *
 * It's guarateed that there is at least one valid solution. If there are many
 * valid solutions you can return any of them.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: [1,1]
 * Explanation: A = 1, B = 1. A + B = n and both A and B don't contain any 0 in
 * their decimal representation.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 11
 * Output: [2,9]
 *
 *
 * Example 3:
 *
 *
 * Input: n = 10000
 * Output: [1,9999]
 *
 *
 * Example 4:
 *
 *
 * Input: n = 69
 * Output: [1,68]
 *
 *
 * Example 5:
 *
 *
 * Input: n = 1010
 * Output: [11,999]
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= n <= 10^4
 *
 */
package main

import "fmt"

func main() {
	num := 1099
	fmt.Printf("%v\n", getNoZeroIntegers(num))
}

// @lc code=start
func isNoZeroInteger(num int) bool {
	for num > 0 {
		mod := num % 10
		if mod == 0 {
			return false
		}
		num = num / 10
	}
	return true
}
func getNoZeroIntegers(n int) []int {
	ans := []int{}

	for i := 1; i <= n/2; i++ {
		if isNoZeroInteger(i) && isNoZeroInteger(n-i) {
			ans = append(ans, i, n-i)
			break
		}
	}

	return ans
}

// @lc code=end
