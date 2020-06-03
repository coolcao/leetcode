/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (45.93%)
 * Likes:    7484
 * Dislikes: 346
 * Total Accepted:    984.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := maxSubArray(nums)
	fmt.Println(max)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// f(i)=max(f(i), f(i)+nums[i])
// @lc code=start
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	curSum := nums[0]
	max := nums[0]

	for i := 1; i < length; i++ {
		if curSum > 0 {
			curSum = curSum + nums[i]
		} else {
			curSum = nums[i]
		}
		max = getMax(max, curSum)
	}
	return max
}

// @lc code=end
