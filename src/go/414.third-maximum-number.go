/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 *
 * https://leetcode.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (30.15%)
 * Likes:    629
 * Dislikes: 1154
 * Total Accepted:    143.8K
 * Total Submissions: 473.7K
 * Testcase Example:  '[3,2,1]'
 *
 * Given a non-empty array of integers, return the third maximum number in this
 * array. If it does not exist, return the maximum number. The time complexity
 * must be in O(n).
 *
 * Example 1:
 *
 * Input: [3, 2, 1]
 *
 * Output: 1
 *
 * Explanation: The third maximum is 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [1, 2]
 *
 * Output: 2
 *
 * Explanation: The third maximum does not exist, so the maximum (2) is
 * returned instead.
 *
 *
 *
 * Example 3:
 *
 * Input: [2, 2, 3, 1]
 *
 * Output: 1
 *
 * Explanation: Note that the third maximum here means the third maximum
 * distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 *
 *
 */

package main

import "fmt"

// @lc code=start
func thirdMax(nums []int) int {
	max, min, second, third := nums[0], nums[0], 0, 0
	for _, n := range nums {
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
	}
	fmt.Printf("max: %d, min: %d\n", max, min)
	second, third = min, min
	for _, n := range nums {
		if second <= n && n < max {
			second = n
		}
	}
	for _, n := range nums {
		if third <= n && n < second {
			third = n
		}
	}
	fmt.Printf("%d, %d, %d\n", max, second, third)
	if second == third {
		return max
	}
	return third
}

// @lc code=end

func main() {
	nums := []int{1, 2, 2, 5, 3, 5}
	thirdMax(nums)
}
