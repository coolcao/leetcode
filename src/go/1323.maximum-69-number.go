/*
 * @lc app=leetcode id=1323 lang=golang
 *
 * [1323] Maximum 69 Number
 *
 * https://leetcode.com/problems/maximum-69-number/description/
 *
 * algorithms
 * Easy (78.09%)
 * Likes:    265
 * Dislikes: 43
 * Total Accepted:    43.2K
 * Total Submissions: 55.5K
 * Testcase Example:  '9669'
 *
 * Given a positive integer num consisting only of digits 6 and 9.
 *
 * Return the maximum number you can get by changing at most one digit (6
 * becomes 9, and 9 becomes 6).
 *
 *
 * Example 1:
 *
 *
 * Input: num = 9669
 * Output: 9969
 * Explanation:
 * Changing the first digit results in 6669.
 * Changing the second digit results in 9969.
 * Changing the third digit results in 9699.
 * Changing the fourth digit results in 9666.
 * The maximum number is 9969.
 *
 *
 * Example 2:
 *
 *
 * Input: num = 9996
 * Output: 9999
 * Explanation: Changing the last digit 6 to 9 results in the maximum number.
 *
 * Example 3:
 *
 *
 * Input: num = 9999
 * Output: 9999
 * Explanation: It is better not to apply any change.
 *
 *
 * Constraints:
 *
 *
 * 1 <= num <= 10^4
 * num's digits are 6 or 9.
 *
 */
package main

import "fmt"

func main() {
	num := 9669
	ans := maximum69Number(num)
	fmt.Printf("ans: %d\n", ans)
}

// @lc code=start
func maximum69Number(num int) int {
	n := num
	ans := 0
	mask := 1000
	for mask > 0 {
		digit := n / mask
		if digit == 6 {
			ans += 9*mask + (n % mask)
			return ans
		}
		ans += digit * mask
		n = n % mask
		mask = mask / 10
	}
	return ans
}

// @lc code=end
