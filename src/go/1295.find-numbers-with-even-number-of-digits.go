/*
 * @lc app=leetcode id=1295 lang=golang
 *
 * [1295] Find Numbers with Even Number of Digits
 *
 * https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/
 *
 * algorithms
 * Easy (84.64%)
 * Likes:    334
 * Dislikes: 52
 * Total Accepted:    122.1K
 * Total Submissions: 148.3K
 * Testcase Example:  '[12,345,2,6,7896]'
 *
 * Given an array nums of integers, return how many of them contain an even
 * number of digits.
 *
 * Example 1:
 *
 *
 * Input: nums = [12,345,2,6,7896]
 * Output: 2
 * Explanation:
 * 12 contains 2 digits (even number of digits).
 * 345 contains 3 digits (odd number of digits).
 * 2 contains 1 digit (odd number of digits).
 * 6 contains 1 digit (odd number of digits).
 * 7896 contains 4 digits (even number of digits).
 * Therefore only 12 and 7896 contain an even number of digits.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [555,901,482,1771]
 * Output: 1
 * Explanation:
 * Only 1771 contains an even number of digits.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 500
 * 1 <= nums[i] <= 10^5
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{100000}
	fmt.Printf("%v\n", findNumbers(nums))
}

// @lc code=start
func isEvenDigitsNumber(num int) bool {
	if (num >= 10 && num <= 99) || (num >= 1000 && num <= 9999) || num == 100000 {
		return true
	}
	return false
}
func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		if isEvenDigitsNumber(n) {
			count++
		}
	}
	return count
}

// @lc code=end
