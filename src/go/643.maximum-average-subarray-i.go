/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 *
 * https://leetcode.com/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (41.00%)
 * Likes:    681
 * Dislikes: 113
 * Total Accepted:    75.8K
 * Total Submissions: 183.4K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 * Given an array consisting of n integers, find the contiguous subarray of
 * given length k that has the maximum average value. And you need to output
 * the maximum average value.
 *
 * Example 1:
 *
 *
 * Input: [1,12,-5,-6,50,3], k = 4
 * Output: 12.75
 * Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
 *
 * Note:
 * 1 <= k <= n <= 30,000.
 * Elements of the given array will be in the range [-10,000, 10,000].
 *
 * 这是一个简单的滑动窗口
 */
package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	ma := findMaxAverage(nums, k)
	fmt.Printf("%v\n", ma)
}

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	length := len(nums)
	sum := 0
	// init
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAverage := float64(sum) / float64(k)
	for i := k; i < length; i++ {
		sum += nums[i]
		sum -= nums[i-k]
		average := float64(sum) / float64(k)
		if maxAverage < average {
			maxAverage = average
		}

	}
	return maxAverage
}

// @lc code=end
