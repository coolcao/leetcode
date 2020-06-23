/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 *
 * https://leetcode.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Easy (19.42%)
 * Likes:    1779
 * Dislikes: 434
 * Total Accepted:    92K
 * Total Submissions: 471.4K
 * Testcase Example:  '[4,2,3]'
 *
 * Given an array nums with n integers, your task is to check if it could
 * become non-decreasing by modifying at most 1 element.
 *
 * We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for
 * every i (0-based) such that (0 <= i <= n - 2).
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [4,2,3]
 * Output: true
 * Explanation: You could modify the first 4 to 1 to get a non-decreasing
 * array.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,2,1]
 * Output: false
 * Explanation: You can't get a non-decreasing array by modify at most one
 * element.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10 ^ 4
 * - 10 ^ 5 <= nums[i] <= 10 ^ 5
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 4, 4, 3, 3}
	check := checkPossibility(nums)
	fmt.Printf("%v\n", check)
}

// @lc code=start
func checkPossibility(nums []int) bool {
	length := len(nums)
	if length <= 2 {
		return true
	}

	tmp := []int{}

	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			tmp = append(tmp, i)
		}
	}
	if len(tmp) == 0 {
		return true
	}
	if len(tmp) >= 2 {
		return false
	}
	// 只有一个非递增位置
	idx := tmp[0]
	fmt.Printf("idx: %d\n", idx)
	if idx == 0 || idx == length-2 {
		return true
	}

	if nums[idx+1] >= nums[idx-1] || nums[idx+2] >= nums[idx] {
		return true
	}

	return false
}

// @lc code=end
