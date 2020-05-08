/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (36.99%)
 * Likes:    1992
 * Dislikes: 98
 * Total Accepted:    247.2K
 * Total Submissions: 663.7K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of n positive integers and a positive integer s, find the
 * minimal length of a contiguous subarray of which the sum ≥ s. If there isn't
 * one, return 0 instead.
 *
 * Example:
 *
 *
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem
 * constraint.
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution of
 * which the time complexity is O(n log n).
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	s := 15
	min := minSubArrayLen(s, nums)
	fmt.Printf("min: %d\n", min)
}

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	length := len(nums)

	min := length + 1

	// 滑动窗口的左右指针
	left, right := 0, 0
	// 窗口内元素的和
	sum := 0
	// 当和小于s时，增大窗口
	for sum < s && right < length {

		// 如果最小窗口长度已经是1，那么窗口可终止滑动
		if min == 1 {
			break
		}

		sum += nums[right]
		right++

		// 当和大于等于s时，缩小窗口
		for sum >= s {
			// 比较此时窗口长度与记录的最小长度
			if min > right-left {
				min = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if min == length+1 {
		return 0
	}
	return min
}

// @lc code=end
