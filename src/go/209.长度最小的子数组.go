/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (44.65%)
 * Likes:    521
 * Dislikes: 0
 * Total Accepted:    105.9K
 * Total Submissions: 237K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续
 * 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
 *
 * 示例：
 * 输入：s = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 * 进阶：
 * 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 */
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 15
	l := minSubArrayLen(s, nums)
	fmt.Println(l)
}

// @lc code=start
// 滑动窗口
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	minLen := length
	left, right := 0, 0
	sum := 0
	for right < length {
		sum += nums[right]
		right++
		// 满足条件
		for sum >= s {
			minLen = min(minLen, right-left)
			sum -= nums[left]
			left++
		}
	}
	if left == 0 && right == length && sum < s {
		return 0
	}
	return minLen
}

// @lc code=end
